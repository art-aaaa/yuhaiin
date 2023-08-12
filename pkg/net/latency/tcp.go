package latency

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Asutorufa/yuhaiin/pkg/net/netapi"
)

func HTTP(p netapi.Proxy, target string) (time.Duration, error) {
	tr := &http.Transport{
		DisableKeepAlives: true,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			ad, err := netapi.ParseAddress(netapi.PaseNetwork(network), addr)
			if err != nil {
				return nil, fmt.Errorf("parse address failed: %w", err)
			}
			return p.Conn(ctx, ad)
		},
	}
	defer tr.CloseIdleConnections()

	start := time.Now()
	resp, err := (&http.Client{Transport: tr}).Get(target)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return time.Since(start), nil
}
