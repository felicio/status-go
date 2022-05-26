package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"go.uber.org/zap"

	"github.com/status-im/status-go/logutils"
)

type Server struct {
	run      bool
	server   *http.Server
	logger   *zap.Logger
	cert     *tls.Certificate
	hostname string
	port     int
	listener net.Listener
}

func NewServer(cert *tls.Certificate, hostname string) Server {
	return Server{logger: logutils.ZapLogger(), cert: cert, hostname: hostname}
}

func (s *Server) setListener(l net.Listener) {
	s.listener = l
}

func (s *Server) resetListener() {
	l := s.logger.Named("resetListener").With(zap.String("name", "resetListener"))
	l.Info("Has been fired")
	s.listener = nil
}

func (s *Server) makeListener(port int) {
	l := s.logger.Named("makeListener").With(zap.String("name", "makeListener"))
	l.Info("Triggered")

	cfg := &tls.Config{Certificates: []tls.Certificate{*s.cert}, ServerName: s.hostname, MinVersion: tls.VersionTLS12}

	addr := fmt.Sprintf("%s:%d", s.hostname, port)

	listener, err := tls.Listen("tcp", addr, cfg)
	l.Info("tls.Listen", zap.Error(err), zap.String("listener.Addr()", listener.Addr().String()))
	if err != nil {
		l.Error("failed to make listener, retrying", zap.Error(err))

		for {
			listener, err = tls.Listen("tcp", addr, cfg)
			if err == nil {
				break
			}
			l.Error("failed to make listener, retrying", zap.Error(err))
		}

		//s.resetListener()
		err = s.Start()
		if err != nil {
			s.logger.Error("server start failed, giving up", zap.Error(err))
		}
		return
	}

	//s.setListener(listener)
}

// getPort depends on the Server.listener to provide a port number, net.Listener should determine the port.
// This is because there is no way to know what ports are available on the host device in advance
func (s *Server) getPort() int {
	if s.listener == nil {
		return 0
	}

	return s.listener.Addr().(*net.TCPAddr).Port
}

func (s *Server) getHost() string {
	l := s.logger.Named("getHost").With(zap.String("name", "getHost"))
	l.Info("Triggered")
	// TODO consider returning an error if s.getPort returns `0`, as this means that the listener is not ready
	return fmt.Sprintf("%s:%d", s.hostname, s.port)
}

func (s *Server) listenAndServe() {
	l := s.logger.Named("listenAndServe").With(zap.String("name", "listenAndServe"))
	l.Info("Triggered")

	cfg := &tls.Config{Certificates: []tls.Certificate{*s.cert}, ServerName: s.hostname, MinVersion: tls.VersionTLS12}

	// in case of restart, we should use the same port as the first start in order not to break existing links
	//s.makeListener()
	addr := fmt.Sprintf("%s:%d", s.hostname, s.port)

	listener, err := tls.Listen("tcp", addr, cfg)
	if err != nil {
		s.logger.Error("failed to start server, retrying", zap.Error(err))
		s.port = 0
		err = s.Start()
		if err != nil {
			s.logger.Error("server start failed, giving up", zap.Error(err))
		}
		return
	}

	s.port = listener.Addr().(*net.TCPAddr).Port
	s.run = true

	err = s.server.Serve(listener)
	if err != http.ErrServerClosed {
		s.logger.Error("server failed unexpectedly, restarting", zap.Error(err))
		err = s.Start()
		if err != nil {
			s.logger.Error("server start failed, giving up", zap.Error(err))
		}
		return
	}

	s.run = false
}

func (s *Server) Start() error {
	l := s.logger.Named("Start()").With(zap.String("name", "Start()"))
	l.Info("has been fired")
	go s.listenAndServe()
	return nil
}

func (s *Server) Stop() error {
	l := s.logger.Named("Stop()").With(zap.String("name", "Stop()"))
	l.Info("has been fired")
	if s.server != nil {
		return s.server.Shutdown(context.Background())
	}

	return nil
}

func (s *Server) ToForeground() {
	l := s.logger.Named("ToForeground").With(zap.String("name", "ToForeground"))
	l.Info("has been fired")
	if !s.run && (s.server != nil) {
		l.Info("in if statement")
		err := s.Start()
		l.Info("s.Start() error", zap.Error(err))
		if err != nil {
			s.logger.Error("server start failed during foreground transition", zap.Error(err))
		}
	}
}

func (s *Server) ToBackground() {
	l := s.logger.Named("ToBackground").With(zap.String("name", "ToBackground"))
	l.Info("has been triggered")
	if s.run {
		l.Info("in if statement")
		err := s.Stop()
		l.Info("s.Stop() error", zap.Error(err))
		if err != nil {
			s.logger.Error("server stop failed during background transition", zap.Error(err))
		}
	}
}

func (s *Server) WithHandlers(handlers HandlerPatternMap) {
	switch {
	case s.server != nil && s.server.Handler != nil:
		break
	case s.server != nil && s.server.Handler == nil:
		s.server.Handler = http.NewServeMux()
	default:
		s.server = &http.Server{}
		s.server.Handler = http.NewServeMux()
	}

	for p, h := range handlers {
		s.server.Handler.(*http.ServeMux).HandleFunc(p, h)
	}
}

func (s *Server) MakeBaseURL() *url.URL {
	return &url.URL{
		Scheme: "https",
		Host:   s.getHost(),
	}
}
