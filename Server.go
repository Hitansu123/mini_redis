package main

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
	}
}
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	fmt.Println("Listening in", s.listenAddr)
	defer ln.Close()
	s.ln = ln
	go s.acceptloop()
	<-s.quitch

	return nil
}
func (s *Server) acceptloop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Accept error", err)
			continue
		}
		go s.readLoop(conn)
	}
}
func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
	}
}
