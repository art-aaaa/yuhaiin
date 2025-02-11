/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2022 WireGuard LLC. All Rights Reserved.
 */

package wireguard

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Asutorufa/yuhaiin/pkg/log"
	"github.com/Asutorufa/yuhaiin/pkg/net/netapi"
	"github.com/Asutorufa/yuhaiin/pkg/protos/node/point"
	"github.com/Asutorufa/yuhaiin/pkg/protos/node/protocol"
	"github.com/tailscale/wireguard-go/device"
	"gvisor.dev/gvisor/pkg/tcpip/adapters/gonet"
)

type Wireguard struct {
	netapi.EmptyDispatch
	net  *Net
	bind *netBindClient

	conf *protocol.Wireguard
	mu   sync.Mutex

	count atomic.Int64

	lastNewConn time.Time
	idleTimeout time.Duration

	device *device.Device
}

func init() {
	point.RegisterProtocol(NewClient)
}

func NewClient(conf *protocol.Protocol_Wireguard) point.WrapProxy {
	return func(p netapi.Proxy) (netapi.Proxy, error) {

		if conf.Wireguard.IdleTimeout == 0 {
			conf.Wireguard.IdleTimeout = 60 * 5
		}
		if conf.Wireguard.IdleTimeout <= 30 {
			conf.Wireguard.IdleTimeout = 30
		}

		return &Wireguard{
			conf:        conf.Wireguard,
			idleTimeout: time.Duration(conf.Wireguard.IdleTimeout) * time.Second,
		}, nil
	}
}

func (w *Wireguard) collect() {
	readyClose := false

	for {
		time.Sleep(w.idleTimeout)

		br := func() bool {
			w.mu.Lock()
			defer w.mu.Unlock()

			log.Debug("wireguard check idle timeout")

			if w.count.Load() > 0 {
				readyClose = false
				return false
			}

			if !w.lastNewConn.IsZero() && time.Since(w.lastNewConn) < time.Minute {
				readyClose = false
				return false
			}

			if readyClose {
				log.Debug("wireguard closing")
				if w.device != nil {
					w.device.Close()
					w.device = nil
				}

				if w.bind != nil {
					w.bind.Close()
					w.bind = nil
				}
				log.Debug("wireguard closed")
				w.net = nil
				return true
			}

			log.Debug("wireguard ready to close")

			readyClose = true
			return false
		}()

		if br {
			break
		}
	}
}

func (w *Wireguard) initNet() (*Net, error) {
	net := w.net
	if net != nil {
		return net, nil
	}

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.net != nil {
		return w.net, nil
	}

	dev, bind, net, err := makeVirtualTun(w.conf)
	if err != nil {
		return nil, err
	}

	w.device = dev
	w.net = net
	w.bind = bind
	go w.collect()

	return net, nil
}

func (w *Wireguard) Conn(ctx context.Context, addr netapi.Address) (net.Conn, error) {
	net, err := w.initNet()
	if err != nil {
		return nil, err
	}

	addrPort := addr.AddrPort(ctx)

	if addrPort.Err != nil {
		return nil, addrPort.Err
	}

	conn, err := net.DialContextTCPAddrPort(ctx, addrPort.V)
	if err != nil {
		return nil, err
	}

	w.count.Add(1)
	w.lastNewConn = time.Now()

	return &wrapGoNetTcpConn{w, conn}, nil
}

type wrapGoNetTcpConn struct {
	wireguard *Wireguard
	*gonet.TCPConn
}

func (w *wrapGoNetTcpConn) Close() error {
	w.wireguard.count.Add(-1)
	return w.TCPConn.Close()
}

func (w *Wireguard) PacketConn(ctx context.Context, addr netapi.Address) (net.PacketConn, error) {
	net, err := w.initNet()
	if err != nil {
		return nil, err
	}

	goUC, err := net.ListenUDP(nil)
	if err != nil {
		return nil, err
	}

	w.count.Add(1)
	w.lastNewConn = time.Now()

	return &wrapGoNetUdpConn{w, goUC}, nil
}

type wrapGoNetUdpConn struct {
	wireguard *Wireguard
	*gonet.UDPConn
}

func (w *wrapGoNetUdpConn) Close() error {
	w.wireguard.count.Add(-1)
	return w.UDPConn.Close()
}

func (w *wrapGoNetUdpConn) WriteTo(buf []byte, addr net.Addr) (int, error) {
	a, err := netapi.ParseSysAddr(addr)
	if err != nil {
		return 0, err
	}

	ur := a.UDPAddr(context.TODO())

	if ur.Err != nil {
		return 0, ur.Err
	}

	return w.UDPConn.WriteTo(buf, ur.V)
}

// creates a tun interface on netstack given a configuration
func makeVirtualTun(h *protocol.Wireguard) (*device.Device, *netBindClient, *Net, error) {
	endpoints, err := parseEndpoints(h)
	if err != nil {
		return nil, nil, nil, err
	}
	tun, tnet, err := CreateNetTUN(endpoints, int(h.Mtu))
	if err != nil {
		return nil, nil, nil, err
	}

	bind := newNetBindClient(h.GetReserved())
	// dev := device.NewDevice(tun, conn.NewDefaultBind(), nil /* device.NewLogger(device.LogLevelVerbose, "") */)
	dev := device.NewDevice(
		tun,
		bind,
		&device.Logger{
			Verbosef: func(format string, args ...any) {
				log.Output(2, slog.LevelDebug, fmt.Sprintf(format, args...))
			},
			Errorf: func(format string, args ...any) {
				log.Output(2, slog.LevelError, fmt.Sprintf(format, args...))
			},
		})

	err = dev.IpcSet(createIPCRequest(h))
	if err != nil {
		dev.Close()
		return nil, nil, nil, err
	}

	err = dev.Up()
	if err != nil {
		dev.Close()
		return nil, nil, nil, err
	}

	return dev, bind, tnet, nil
}

func base64ToHex(s string) string {
	data, _ := base64.StdEncoding.DecodeString(s)
	return hex.EncodeToString(data)
}

// serialize the config into an IPC request
func createIPCRequest(conf *protocol.Wireguard) string {
	var request bytes.Buffer

	request.WriteString(fmt.Sprintf("private_key=%s\n", base64ToHex(conf.SecretKey)))

	for _, peer := range conf.Peers {
		request.WriteString(fmt.Sprintf("public_key=%s\nendpoint=%s\n", base64ToHex(peer.PublicKey), peer.Endpoint))
		if peer.KeepAlive != 0 {
			request.WriteString(fmt.Sprintf("persistent_keepalive_interval=%d\n", peer.KeepAlive))
		}
		if peer.PreSharedKey != "" {
			request.WriteString(fmt.Sprintf("preshared_key=%s\n", base64ToHex(peer.PreSharedKey)))
		}

		for _, ip := range peer.AllowedIps {
			request.WriteString(fmt.Sprintf("allowed_ip=%s\n", ip))
		}
	}

	return request.String()[:request.Len()]
}
