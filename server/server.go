package server

import (
	"fmt"
	"net"

	"github.com/Toffee-iZt/gomine/log"
	"github.com/Toffee-iZt/gomine/server/sync"
)

// Server ...
type Server struct {
	command chan sync.Command
	logger  *log.Logger
	host    string
	port    int
}

// New creates new server.
func New(cfg Config) *Server {
	s := Server{
		command: make(chan sync.Command, 1),
		logger:  log.New("SERVER", nil, log.FullLog, true),
		host:    cfg.Host,
		port:    cfg.Port,
	}
	return &s
}

// Start ...
func (s *Server) Start() error {
	tcpAddr := fmt.Sprintf("%s:%d", s.host, s.port)
	addr, err := net.ResolveTCPAddr("tcp", tcpAddr)
	if err != nil {
		return err
	}

	tcp, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	s.logger.Info("Server started on %s", tcpAddr)

	for {
		var conn *net.TCPConn
		conn, err = tcp.AcceptTCP()
		if err != nil {
			break
		}

		conn.SetKeepAlive(true)

		go s.handle(conn)
	}

	tcp.Close()
	return err
}

func (s *Server) handle(conn *net.TCPConn) {}
