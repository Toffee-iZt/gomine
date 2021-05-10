// Package net pack network connection for Minecraft.
package net

import (
	"bufio"
	"crypto/cipher"
	"io"
	"net"
	"time"

	"github.com/Toffee-iZt/gomine/proto/net/packet"
)

// Listener is a minecraft Listener
type Listener struct {
	Listener *net.TCPListener
}

// ListenMC listen as TCP but Accept a mc Conn
func ListenMC(addr string) (*Listener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return &Listener{l}, nil
}

// Accept waits for and returns minecraft connection
func (l Listener) Accept() (Conn, error) {
	conn, err := l.Listener.AcceptTCP()
	return Conn{
		Socket: conn,
		Reader: bufio.NewReader(conn),
		Writer: conn,
	}, err
}

// Conn is a minecraft connection
type Conn struct {
	Socket net.Conn
	io.Reader
	io.Writer

	threshold int
}

// DialMC creates a Minecraft connection
func DialMC(addr string) (*Conn, error) {
	conn, err := net.Dial("tcp", addr)
	return WrapConn(conn), err
}

// DialMCTimeout acts like DialMC but takes a timeout.
func DialMCTimeout(addr string, timeout time.Duration) (*Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	return WrapConn(conn), err
}

// WrapConn warps an net.Conn to minecraft Conn
func WrapConn(conn net.Conn) *Conn {
	return &Conn{
		Socket: conn,
		Reader: conn,
		Writer: conn,
	}
}

//Close closes the connection
func (c *Conn) Close() error {
	return c.Socket.Close()
}

// ReadPacket reads a Packet from Conn
func (c *Conn) ReadPacket(p *packet.Packet) error {
	return p.UnPack(c.Reader, c.threshold)
}

//WritePacket writes a Packet to Conn
func (c *Conn) WritePacket(p packet.Packet) error {
	return p.Pack(c, c.threshold)
}

// SetCipher loads the decode/encode stream to this Conn
func (c *Conn) SetCipher(ecoStream, decoStream cipher.Stream) {
	c.Reader = cipher.StreamReader{
		S: decoStream,
		R: c.Socket,
	}
	c.Writer = cipher.StreamWriter{
		S: ecoStream,
		W: c.Socket,
	}
}

// SetThreshold sets threshold to Conn.
// The data packet with length longer then threshold
// will be compress when sending.
func (c *Conn) SetThreshold(t int) {
	c.threshold = t
}
