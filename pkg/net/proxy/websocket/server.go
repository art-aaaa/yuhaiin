package websocket

import (
	"errors"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/Asutorufa/yuhaiin/pkg/log"
	"github.com/Asutorufa/yuhaiin/pkg/net/proxy/websocket/websocket"
)

type Server struct {
	net.Listener
	server   *http.Server
	connChan chan *Connection
	closed   bool
	lock     sync.RWMutex
}

func NewServer(lis net.Listener) *Server {
	s := &Server{
		Listener: lis,
		connChan: make(chan *Connection, 20),
	}
	s.server = &http.Server{Handler: s}

	go func() {
		defer s.Close()
		s.server.Serve(lis)
	}()

	return s
}

func (s *Server) Close() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.closed {
		return nil
	}

	err := s.server.Close()
	if er := s.Listener.Close(); er != nil {
		err = errors.Join(err, er)
	}
	close(s.connChan)
	s.closed = true
	return err
}

func (s *Server) Accept() (net.Conn, error) {
	conn, ok := <-s.connChan
	if !ok {
		return nil, net.ErrClosed
	}

	return conn, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.closed {
		return
	}

	if strings.ToLower(req.Header.Get("Upgrade")) != "websocket" ||
		!strings.Contains(strings.ToLower(req.Header.Get("Connection")), "upgrade") {
		return
	}

	conn, buf, err := w.(http.Hijacker).Hijack()
	if err != nil {
		log.Errorln("hijack failed:", err)
		return
	}

	wsconn, err := websocket.NewServerConn(conn, buf, req, &websocket.Config{}, nil)
	if err != nil {
		log.Errorln("new websocket server conn failed:", err)
		return
	}

	s.connChan <- &Connection{wsconn, conn}
}
