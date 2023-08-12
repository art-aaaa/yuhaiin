package shadowsocks

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/Asutorufa/yuhaiin/pkg/net/netapi"
	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/shadowsocks/core"
	s5c "github.com/Asutorufa/yuhaiin/pkg/net/proxy/socks5/client"
	"github.com/Asutorufa/yuhaiin/pkg/protos/node/protocol"
	"github.com/Asutorufa/yuhaiin/pkg/protos/statistic"
)

// Shadowsocks shadowsocks
type Shadowsocks struct {
	cipher core.Cipher
	p      netapi.Proxy
	netapi.EmptyDispatch
}

func New(config *protocol.Protocol_Shadowsocks) protocol.WrapProxy {
	c := config.Shadowsocks
	return func(p netapi.Proxy) (netapi.Proxy, error) {
		cipher, err := core.PickCipher(strings.ToUpper(c.Method), nil, c.Password)
		if err != nil {
			return nil, err
		}

		return &Shadowsocks{cipher: cipher, p: p}, nil
	}
}

// Conn .
func (s *Shadowsocks) Conn(ctx context.Context, addr netapi.Address) (conn net.Conn, err error) {
	conn, err = s.p.Conn(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("dial to %s failed: %w", addr, err)
	}

	if x, ok := conn.(*net.TCPConn); ok {
		_ = x.SetKeepAlive(true)
	}

	conn = s.cipher.StreamConn(conn)
	if _, err = conn.Write(s5c.ParseAddr(addr)); err != nil {
		conn.Close()
		return nil, fmt.Errorf("shadowsocks write target failed: %w", err)
	}
	return conn, nil
}

// PacketConn .
func (s *Shadowsocks) PacketConn(ctx context.Context, tar netapi.Address) (net.PacketConn, error) {
	pc, err := s.p.PacketConn(ctx, tar)
	if err != nil {
		return nil, fmt.Errorf("create packet conn failed")
	}

	return NewPacketConn(s.cipher.PacketConn(pc)), nil
}

type packetConn struct{ net.PacketConn }

func NewPacketConn(conn net.PacketConn) net.PacketConn { return &packetConn{conn} }

func (v *packetConn) ReadFrom(b []byte) (int, net.Addr, error) {
	n, _, err := v.PacketConn.ReadFrom(b)
	if err != nil {
		return 0, nil, fmt.Errorf("read udp from shadowsocks failed: %w", err)
	}

	addr, err := s5c.ResolveAddr(bytes.NewBuffer(b[:n]))
	if err != nil {
		return 0, nil, fmt.Errorf("resolve address failed: %w", err)
	}

	return copy(b, b[len(addr):n]), addr.Address(statistic.Type_udp), nil
}

func (v *packetConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	ad, err := netapi.ParseSysAddr(addr)
	if err != nil {
		return 0, err
	}
	return v.PacketConn.WriteTo(bytes.Join([][]byte{s5c.ParseAddr(ad), b}, nil), addr)
}
