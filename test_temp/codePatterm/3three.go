package codePatterm

import (
	"crypto/tls"
	"fmt"
	"time"
)

//FUNCTIONAL OPTIONS
//比如配置的处理，一些 必填参数+可选参数
//1.定义一个函数类型
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}
type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}

func NewServer(addr string, prot int, options ...Option) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     prot,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}
func TestFuncOption() {
	s1, _ := NewServer("localhost", 1010)
	s2, _ := NewServer("localhost", 1024, Protocol("udp"))
	s3, _ := NewServer("127.0.0.1", 8080, Timeout(60*time.Second), MaxConns(2000))

	fmt.Printf("s1:%v,s2:%v,s3:%v", s1, s2, s3)
}
