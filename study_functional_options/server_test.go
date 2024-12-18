package study_functional_options

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	s1, _ := NewServer("localhost", 8080)

	if s1.Addr != "localhost" {
		t.Errorf("expected address to be localhost, got %s", s1.Addr)
	}

	if s1.Port != 8080 {
		t.Errorf("expected port to be 8080, got %d", s1.Port)
	}

	if s1.Protocol != "tcp" {
		t.Errorf("expected protocol to be tcp, got %s", s1.Protocol)
	}

	if s1.Timeout != 30*time.Second {
		t.Errorf("expected timeout to be 30, got %d", s1.Timeout)
	}

	s2, _ := NewServer("localhost", 8080, Protocol("udp"))

	if s2.Protocol != "udp" {
		t.Errorf("expected protocol to be udp, got %s", s2.Protocol)
	}

	s3, _ := NewServer("localhost", 8080, Timeout(10*time.Second), MaxConns(100))

	if s3.Timeout != 10*time.Second {
		t.Errorf("expected timeout to be 10, got %d", s3.Timeout)
	}

	if s3.MaxConns != 100 {
		t.Errorf("expected max connections to be 100, got %d", s3.MaxConns)
	}

}
