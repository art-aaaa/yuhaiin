package sniffy

import (
	"net"

	"github.com/Asutorufa/yuhaiin/pkg/net/netapi"
	"github.com/Asutorufa/yuhaiin/pkg/net/sniffy/bittorrent"
	"github.com/Asutorufa/yuhaiin/pkg/protos/config/bypass"
	"github.com/Asutorufa/yuhaiin/pkg/utils/pool"
)

type entry[T any] struct {
	name    string
	checker func([]byte) (T, bool)
}

type Sniffier[T any] struct {
	streamChecker []entry[T]
	packetChecker []entry[T]
}

func New() *Sniffier[bypass.Mode] {
	return &Sniffier[bypass.Mode]{
		streamChecker: []entry[bypass.Mode]{
			{
				name: "bittorrent",
				checker: func(b []byte) (bypass.Mode, bool) {
					_, err := bittorrent.SniffBittorrent(b)
					if err == nil {
						return bypass.Mode_direct, true
					}

					return bypass.Mode_bypass, false
				},
			},
		},

		packetChecker: []entry[bypass.Mode]{
			{
				name: "bittorrent_utp",
				checker: func(b []byte) (bypass.Mode, bool) {
					_, err := bittorrent.SniffUTP(b)
					if err == nil {
						return bypass.Mode_direct, true
					}

					return bypass.Mode_bypass, false
				},
			},
		},
	}
}

func (s *Sniffier[T]) Packet(b []byte) (T, string, bool) {
	for _, c := range s.packetChecker {
		t, ok := c.checker(b)
		if ok {
			return t, c.name, ok
		}
	}

	var t T
	return t, "", false
}

func (s *Sniffier[T]) Stream(c net.Conn) (net.Conn, T, string, bool) {
	buf := pool.GetBytesV2(pool.DefaultSize)

	n, _ := c.Read(buf.Bytes())

	var t T
	if n <= 0 {
		pool.PutBytesV2(buf)
		return c, t, "", false
	}

	buf.ResetSize(0, n)

	c = newConn(c, buf)

	for _, ck := range s.streamChecker {
		t, ok := ck.checker(buf.Bytes())
		if ok {
			return c, t, ck.name, ok
		}
	}

	return c, t, "", false
}

type conn struct {
	net.Conn

	buf *pool.Bytes
}

func newConn(c net.Conn, buf *pool.Bytes) net.Conn {
	return &conn{
		Conn: netapi.NewPrefixBytesConn(c, buf.Bytes()),
		buf:  buf,
	}
}

func (c *conn) Close() error {
	pool.PutBytesV2(c.buf)
	return c.Conn.Close()
}
