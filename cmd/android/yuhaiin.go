package yuhaiin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/Asutorufa/yuhaiin/internal/app"
	"github.com/Asutorufa/yuhaiin/pkg/log"
	"github.com/Asutorufa/yuhaiin/pkg/net/dialer"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type App struct {
	lis *http.Server

	mu      sync.Mutex
	started atomic.Bool
}

func (a *App) Start(opt *Opts) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.started.Load() {
		return errors.New("yuhaiin is already running")
	}

	errChan := make(chan error)

	go func() {
		defer a.started.Store(false)

		dialer.DefaultMarkSymbol = opt.TUN.SocketProtect.Protect

		err := app.Start(
			app.StartOpt{
				ConfigPath:    opt.Savepath,
				Setting:       fakeSetting(opt, app.PathGenerator.Config(opt.Savepath)),
				Host:          opt.Host,
				ProcessDumper: NewUidDumper(opt.TUN.UidDumper),
			})
		if err != nil {
			errChan <- err
			return
		}
		defer app.Close()

		lis := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug("http request", "host", r.Host, "method", r.Method, "path", r.URL.Path)
			app.App.Mux.ServeHTTP(w, r)
		})}
		defer lis.Close()

		a.lis = lis
		a.started.Store(true)

		close(errChan)
		defer opt.CloseFallback.Close()

		a.lis.Serve(app.App.HttpListener)
	}()

	return <-errChan
}

func (a *App) Stop() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.Running() {
		return nil
	}

	if a.lis != nil {
		err := a.lis.Close()
		if err != nil {
			return err
		}
	}

	for a.Running() {
		runtime.Gosched()
	}

	return nil
}

func (a *App) Running() bool { return a.started.Load() }

func (a *App) SaveNewBypass(link string) error {
	if !a.Running() || app.App.Tools == nil {
		return fmt.Errorf("proxy service is not start")
	}

	_, err := app.App.Tools.SaveRemoteBypassFile(context.TODO(), &wrapperspb.StringValue{Value: link})
	return err
}
