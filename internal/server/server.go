package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/http2"

	"github.com/sorohimm/misc/log"
)

type Config struct {
	Timeout *Timeout
	Host    string
	Port    int
	TLS     *TLS
}

type Timeout struct {
	Idle       time.Duration
	Read       time.Duration
	Write      time.Duration
	MustShutIn time.Duration
}

type TLS struct {
	Cert string
	Key  string
}

func PrepareHTTP(ctx context.Context, s *HTTPServer, httpAddr string) (func() error, func(error)) {
	logger := log.FromContext(ctx).Sugar()
	exec := func() error {
		logger.With("listen http ", httpAddr).Info("starting http ...")
		return s.Start()
	}

	interr := func(err error) {
		if err != nil {
			logger.Errorf("inerrupt http: %v", err)
		}
		if err = s.Shutdown(); err != nil {
			logger.Errorf("http shutdown: %v", err)
		}
	}

	return exec, interr
}

func NewHTTPServer(ctx context.Context, c *Config, handler http.Handler) (*HTTPServer, error) {
	if c == nil {
		return nil, fmt.Errorf("empty config")
	}

	if c.Host == "" {
		return nil, fmt.Errorf("empty host")
	}

	srv := &http.Server{
		IdleTimeout:  c.Timeout.Idle,
		ReadTimeout:  c.Timeout.Read,
		WriteTimeout: c.Timeout.Write,
		Addr:         net.JoinHostPort(c.Host, strconv.Itoa(c.Port)),
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		Handler:      handler,
	}

	if err := http2.ConfigureServer(srv, &http2.Server{
		IdleTimeout: c.Timeout.Idle,
	}); err != nil {
		return nil, err
	}

	return &HTTPServer{
		Server:          srv,
		cert:            c.TLS.Cert,
		key:             c.TLS.Key,
		shutdownTimeout: c.Timeout.MustShutIn,
	}, nil
}

type HTTPServer struct {
	*http.Server
	cert, key       string
	shutdownTimeout time.Duration
}

func (o *HTTPServer) Start() error {
	if o.cert != "" && o.key != "" {
		return o.Server.ListenAndServeTLS(o.cert, o.key)
	}
	return o.Server.ListenAndServe()
}

func (o *HTTPServer) Shutdown() error {
	stopCtx, cancel := context.WithTimeout(context.Background(), o.shutdownTimeout)
	defer cancel()
	return o.Server.Shutdown(stopCtx)
}
