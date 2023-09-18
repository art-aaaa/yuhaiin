package sysproxy

import (
	"net"
	"net/netip"
	"strconv"

	cb "github.com/Asutorufa/yuhaiin/pkg/protos/config"
	"github.com/Asutorufa/yuhaiin/pkg/protos/config/listener"
	"google.golang.org/protobuf/proto"
)

var server *listener.InboundConfig

func Update(path string) func(s *cb.Setting) {
	return func(s *cb.Setting) {
		if proto.Equal(server, s.Server) {
			return
		}
		UnsetSysProxy(path)
		var http, socks5 string

		for _, v := range s.Server.Servers {
			if s.SystemProxy.Http {
				if v.GetEnabled() && v.GetHttp() != nil {
					http = v.GetHttp().GetHost()
				}

				if v.GetEnabled() && v.GetMix() != nil {
					http = v.GetMix().GetHost()
				}
			}

			if s.SystemProxy.Socks5 {
				if v.GetEnabled() && v.GetSocks5() != nil {
					socks5 = v.GetSocks5().GetHost()
				}

				if v.GetEnabled() && v.GetMix() != nil {
					socks5 = v.GetMix().GetHost()
				}
			}
		}

		SetSysProxy(path, replaceUnspecified(http), replaceUnspecified(socks5))
		server = s.Server
	}
}

func replaceUnspecified(s string) string {
	if ip, err := netip.ParseAddrPort(s); err == nil {
		if ip.Addr().IsUnspecified() {
			if ip.Addr().Is6() {
				return net.JoinHostPort(net.IPv6loopback.String(), strconv.Itoa(int(ip.Port())))
			} else {
				return net.JoinHostPort("127.0.0.1", strconv.Itoa(int(ip.Port())))

			}
		}
	}

	return s
}

func Unset(path string) {
	UnsetSysProxy(path)
}
