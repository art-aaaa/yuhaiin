package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Asutorufa/yuhaiin/internal/app"
	"github.com/Asutorufa/yuhaiin/internal/appapi"
	"github.com/Asutorufa/yuhaiin/internal/version"
	"github.com/Asutorufa/yuhaiin/pkg/components/config"
	"github.com/Asutorufa/yuhaiin/pkg/log"
	"github.com/Asutorufa/yuhaiin/pkg/net/netapi"
	pc "github.com/Asutorufa/yuhaiin/pkg/protos/config"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

var install = func(args []string) error { panic("not implement") }
var uninstall = func(args []string) error { panic("not implement") }
var restart = func(args []string) error {
	if err := stop(args); err != nil {
		return err
	}
	return start(args)
}
var stop = func(args []string) error { panic("not implement") }
var start = func(args []string) error { panic("not implement") }
var showVersion = func(args []string) error { fmt.Print(version.String()); return nil }

var subCommand = map[string]*func(args []string) error{
	"install":   &install,
	"uninstall": &uninstall,
	"restart":   &restart,
	"version":   &showVersion,
	"-v":        &showVersion,
	"start":     &start,
	"stop":      &stop,
}

var processDumper netapi.ProcessDumper

func main() {
	host := flag.String("host", "0.0.0.0:50051", "gRPC and http listen host")
	path := flag.String("path", pc.DefaultConfigDir(), "save data path")
	webdir := flag.String("eweb", "", "external web page")
	flag.Parse()

	if len(os.Args) > 1 {
		if x, ok := subCommand[strings.ToLower(os.Args[1])]; ok {
			var args []string
			for _, v := range os.Args[1:] {
				if v == "install" || v == "uninstall" || v == "restart" {
					continue
				}

				args = append(args, v)
			}

			if err := (*x)(args); err != nil {
				log.Error(err.Error())
				panic(err)
			}
			return
		}
	}

	if *webdir != "" && os.Getenv("EXTERNAL_WEB") == "" {
		os.Setenv("EXTERNAL_WEB", *webdir)
	}

	setting := config.NewConfig(app.PathGenerator.Config(*path))
	grpcserver := grpc.NewServer()

	app, err := app.Start(appapi.Start{
		ConfigPath:    *path,
		Host:          *host,
		Setting:       setting,
		GRPCServer:    grpcserver,
		ProcessDumper: processDumper,
	})
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer app.Close()

	// _ = os.Rename(filepath.Join(*savepath, "pbo.pprof"),
	// 	filepath.Join(*savepath, fmt.Sprintf("pbo_%d.pprof", time.Now().Unix())))
	// f, err := os.Create(filepath.Join(*savepath, "pbo.pprof"))
	// if err == nil {
	// 	defer f.Close() // error handling omitted for example
	// 	// runtime.SetCPUProfileRate(100)
	// 	if err := pprof.StartCPUProfile(f); err == nil {
	// 		log.Debug("start pprof")
	// 		defer pprof.StopCPUProfile()
	// 	} else {
	// 		f.Close()
	// 		log.Error(err.Error())
	// 	}
	// } else {
	// 	log.Error(err.Error())
	// }

	errChan := make(chan error)

	go func() {
		// h2c for grpc insecure mode
		errChan <- http.Serve(app.HttpListener, h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Debug("http request", "host", r.Host, "method", r.Method, "path", r.URL.Path)

			if grpcserver != nil && r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcserver.ServeHTTP(w, r)
			} else {
				app.Mux.ServeHTTP(w, r)
			}
		}), &http2.Server{}))
	}()

	// listen system signal
	signChannel := make(chan os.Signal, 1)
	signal.Notify(signChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	run(app, errChan, signChannel)
}

var run = func(app *appapi.Components, errChan chan error, signChannel chan os.Signal) {
	select {
	case err := <-errChan:
		log.Error("http server error", "err", err)
	case <-signChannel:
		if app.HttpListener != nil {
			app.HttpListener.Close()
		}
	}
}
