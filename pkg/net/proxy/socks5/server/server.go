package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/Asutorufa/yuhaiin/pkg/log"
	"github.com/Asutorufa/yuhaiin/pkg/net/dialer"
	"github.com/Asutorufa/yuhaiin/pkg/net/netapi"
	s5c "github.com/Asutorufa/yuhaiin/pkg/net/proxy/socks5/client"
	"github.com/Asutorufa/yuhaiin/pkg/protos/config/listener"
	"github.com/Asutorufa/yuhaiin/pkg/protos/statistic"
	"github.com/Asutorufa/yuhaiin/pkg/utils/pool"
	"github.com/Asutorufa/yuhaiin/pkg/utils/relay"
)

func (s *Socks5) newTCPServer(lis net.Listener) {
	s.lis = lis

	go func() {
		defer s.Close()
		for {
			conn, err := lis.Accept()
			if err != nil {
				log.Error("socks5 accept failed", "err", err)

				if ne, ok := err.(net.Error); ok && ne.Temporary() {
					continue
				}
				return
			}

			go func() {
				if err := s.handle(conn); err != nil {
					if errors.Is(err, netapi.ErrBlocked) {
						log.Debug(err.Error())
					} else {
						log.Error("socks5 server handle failed", "err", err)
					}
				}
			}()

		}
	}()
}

func (s *Socks5) handle(client net.Conn) (err error) {
	b := pool.GetBytes(pool.DefaultSize)
	defer pool.PutBytes(b)

	err = handshake1(client, s.username, s.password, b)
	if err != nil {
		return fmt.Errorf("first hand failed: %w", err)
	}

	if err = handshake2(client, s.handler, b); err != nil {
		return fmt.Errorf("second hand failed: %w", err)
	}

	return
}

func handshake1(client net.Conn, user, key string, buf []byte) error {
	//socks5 first handshake
	if _, err := io.ReadFull(client, buf[:2]); err != nil {
		return fmt.Errorf("read first handshake failed: %w", err)
	}

	if buf[0] != 0x05 { // ver
		writeHandshake1(client, s5c.NoAcceptableMethods)
		return fmt.Errorf("no acceptable method: %d", buf[0])
	}

	nMethods := int(buf[1])

	if nMethods > len(buf) {
		writeHandshake1(client, s5c.NoAcceptableMethods)
		return fmt.Errorf("nMethods length of methods out of buf")
	}

	if _, err := io.ReadFull(client, buf[:nMethods]); err != nil {
		return fmt.Errorf("read methods failed: %w", err)
	}

	needVerify := user != "" || key != ""

	for _, v := range buf[:nMethods] { // range all supported methods
		if !needVerify && v == s5c.NoAuthenticationRequired {
			return writeHandshake1(client, s5c.NoAuthenticationRequired)
		}

		if needVerify && v == s5c.UserAndPassword {
			return verifyUserPass(client, user, key)
		}
	}

	writeHandshake1(client, s5c.NoAcceptableMethods)

	return fmt.Errorf("no acceptable authentication methods: [length: %d, method:%v]", nMethods, buf[:nMethods])
}

func verifyUserPass(client net.Conn, user, key string) error {
	b := pool.GetBytes(pool.DefaultSize)
	defer pool.PutBytes(b)
	// get username and password
	_, err := client.Read(b[:])
	if err != nil {
		return err
	}
	username := b[2 : 2+b[1]]
	password := b[3+b[1] : 3+b[1]+b[2+b[1]]]
	if user != string(username) || key != string(password) {
		writeHandshake1(client, 0x01)
		return fmt.Errorf("verify username and password failed")
	}
	writeHandshake1(client, 0x00)
	return nil
}

func handshake2(client net.Conn, f netapi.Handler, buf []byte) error {
	// socks5 second handshake
	if _, err := io.ReadFull(client, buf[:3]); err != nil {
		return fmt.Errorf("read second handshake failed: %w", err)
	}

	if buf[0] != 0x05 { // ver
		writeHandshake2(client, s5c.NoAcceptableMethods, netapi.EmptyAddr)
		return fmt.Errorf("no acceptable method: %d", buf[0])
	}

	var err error

	switch s5c.CMD(buf[1]) { // mode
	case s5c.Connect:
		var adr s5c.ADDR
		adr, err = s5c.ResolveAddr(client)
		if err != nil {
			return fmt.Errorf("resolve addr failed: %w", err)
		}

		addr := adr.Address(statistic.Type_tcp)

		caddr, err := netapi.ParseSysAddr(client.LocalAddr())
		if err != nil {
			return fmt.Errorf("parse local addr failed: %w", err)
		}
		writeHandshake2(client, s5c.Succeeded, caddr) // response to connect successful

		f.Stream(context.TODO(), &netapi.StreamMeta{
			Source:      client.RemoteAddr(),
			Destination: addr,
			Inbound:     client.LocalAddr(),
			Src:         client,
			Address:     addr,
		})

	case s5c.Udp: // udp
		err = handleUDP(client)

	case s5c.Bind: // bind request
		fallthrough

	default:
		writeHandshake2(client, s5c.CommandNotSupport, netapi.EmptyAddr)
		return fmt.Errorf("not Support Method %d", buf[1])
	}

	if err != nil {
		writeHandshake2(client, s5c.HostUnreachable, netapi.EmptyAddr)
	}
	return err
}

func handleUDP(client net.Conn) error {
	laddr, err := netapi.ParseSysAddr(client.LocalAddr())
	if err != nil {
		return fmt.Errorf("parse sys addr failed: %w", err)
	}
	writeHandshake2(client, s5c.Succeeded, netapi.ParseAddressPort(statistic.Type_tcp, "0.0.0.0", laddr.Port()))
	relay.Copy(io.Discard, client)
	return nil
}

func writeHandshake1(conn net.Conn, errREP byte) error {
	_, err := conn.Write([]byte{0x05, errREP})
	return err
}

func writeHandshake2(conn net.Conn, errREP byte, addr netapi.Address) error {
	_, err := conn.Write(append([]byte{0x05, errREP, 0x00}, s5c.ParseAddr(addr)...))
	return err
}

type Socks5 struct {
	udpServer *udpServer
	lis       net.Listener

	handler  netapi.Handler
	addr     string
	username string
	password string
}

func (s *Socks5) Close() error {
	var err error

	if s.udpServer != nil {
		if er := s.udpServer.Close(); er != nil {
			err = errors.Join(err, er)
		}
	}

	if s.lis != nil {
		if er := s.lis.Close(); er != nil {
			err = errors.Join(err, er)
		}
	}

	return err
}

func NewServerWithListener(lis net.Listener, o *listener.Opts[*listener.Protocol_Socks5]) (netapi.Server, error) {
	s := &Socks5{
		handler:  o.Handler,
		addr:     o.Protocol.Socks5.Host,
		username: o.Protocol.Socks5.Username,
		password: o.Protocol.Socks5.Password,
	}

	err := s.newUDPServer(o.Handler)
	if err != nil {
		s.Close()
		return nil, fmt.Errorf("new udp server failed: %w", err)
	}
	s.newTCPServer(lis)

	return s, nil
}

func NewServer(o *listener.Opts[*listener.Protocol_Socks5]) (netapi.Server, error) {
	lis, err := dialer.ListenContext(context.TODO(), "tcp", o.Protocol.Socks5.Host)
	if err != nil {
		return nil, err
	}

	return NewServerWithListener(lis, o)
}
