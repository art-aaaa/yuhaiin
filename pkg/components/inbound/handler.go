package inbound

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Asutorufa/yuhaiin/pkg/log"
	"github.com/Asutorufa/yuhaiin/pkg/net/nat"
	"github.com/Asutorufa/yuhaiin/pkg/net/netapi"
	"github.com/Asutorufa/yuhaiin/pkg/utils/relay"
)

var Timeout = time.Second * 20

type handler struct {
	dialer     netapi.Proxy
	dnsHandler netapi.DNSHandler
	table      *nat.Table
}

func NewHandler(dialer netapi.Proxy, dnsHandler netapi.DNSHandler) *handler {
	h := &handler{
		dialer:     dialer,
		table:      nat.NewTable(dialer),
		dnsHandler: dnsHandler,
	}

	return h
}

func (s *handler) Stream(ctx context.Context, meta *netapi.StreamMeta) {
	go func() {
		if err := s.stream(ctx, meta); err != nil {
			if errors.Is(err, netapi.ErrBlocked) {
				log.Debug("blocked", "msg", err)
			} else {
				log.Error("stream", "error", err)
			}
		}
	}()
}

func (s *handler) stream(ctx context.Context, meta *netapi.StreamMeta) error {
	ctx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	ctx = netapi.NewStore(ctx)
	defer meta.Src.Close()

	dst := meta.Address
	store := netapi.StoreFromContext(ctx)

	store.Add(netapi.SourceKey{}, meta.Source).
		Add(netapi.DestinationKey{}, meta.Destination)
	if meta.Inbound != nil {
		store.Add(netapi.InboundKey{}, meta.Inbound)
	}

	remote, err := s.dialer.Conn(ctx, dst)
	if err != nil {
		return fmt.Errorf("dial %s failed: %w", dst, err)
	}
	defer remote.Close()

	relay.Relay(meta.Src, remote)
	return nil
}

func (s *handler) Packet(ctx context.Context, pack *netapi.Packet) {
	go func() {
		ctx, cancel := context.WithTimeout(ctx, Timeout)
		defer cancel()

		// TODO hijacking dns

		ctx = netapi.NewStore(ctx)

		if err := s.table.Write(ctx, pack); err != nil {
			log.Error("packet", "error", err)
		}
	}()
}

func (s *handler) Close() error { return s.table.Close() }
