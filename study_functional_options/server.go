package study_functional_options

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

type Options func(*Server)

func Protocol(p string) Options {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(t time.Duration) Options {
	return func(s *Server) {
		s.Timeout = t
	}
}

func MaxConns(m int) Options {
	return func(s *Server) {
		s.MaxConns = m
	}
}

func TLSConfig(t *tls.Config) Options {
	return func(s *Server) {
		s.TLS = t
	}
}

func NewServer(addr string, port int, opts ...func(*Server)) (*Server, error) {
	s := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 10,
		TLS:      nil,
	}

	for _, opt := range opts {
		opt(&s)
	}

	return &s, nil
}
