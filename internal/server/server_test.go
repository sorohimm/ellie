package server

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestPrepareHTTP(t *testing.T) {
	ctx := context.Background()
	s := &HTTPServer{Server: &http.Server{}}
	httpAddr := "localhost:8080"

	exec, interr := PrepareHTTP(ctx, s, httpAddr)
	if exec == nil {
		t.Error("exec func should not be nil")
	}
	if interr == nil {
		t.Error("interr func should not be nil")
	}

	interr(errors.New(""))
}

func TestNewHTTPServer(t *testing.T) {
	// Test case 1: success case
	c := &Config{
		Timeout: &Timeout{
			Idle:       time.Second,
			Read:       time.Second,
			Write:      time.Second,
			MustShutIn: time.Second,
		},
		Host: "127.0.0.1",
		Port: 8080,
		TLS: &TLS{
			Cert: "",
			Key:  "",
		},
	}
	s, err := NewHTTPServer(context.Background(), c, http.NotFoundHandler())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if s == nil {
		t.Error("server should not be nil")
	}

	// Test case 2: empty config
	c = nil
	s, err = NewHTTPServer(context.Background(), c, http.NotFoundHandler())
	if err == nil {
		t.Error("error should not be nil")
	}
	if s != nil {
		t.Error("server should be nil")
	}

	// Test case 3: invalid host
	c = &Config{
		Timeout: &Timeout{
			Idle:       time.Second,
			Read:       time.Second,
			Write:      time.Second,
			MustShutIn: time.Second,
		},
		Host: "",
		Port: 8080,
		TLS: &TLS{
			Cert: "",
			Key:  "",
		},
	}
	s, err = NewHTTPServer(context.Background(), c, http.NotFoundHandler())
	if err == nil {
		t.Error("error should not be nil")
	}
	if s != nil {
		t.Error("server should be nil")
	}
}

func TestHTTPServer_Shutdown(t *testing.T) {
	s := &HTTPServer{Server: &http.Server{Addr: "localhost:8080"}}

	err := s.Shutdown()
	if err != nil {
		t.Errorf("error shutting down server: %v", err)
	}
}
