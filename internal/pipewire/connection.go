package pipewire

import (
	"io"
	"net"
)

type PipewireWriter interface {
	io.Writer

	Seq() uint32
}

type Connection struct {
	Conn net.Conn
	seq  uint32
}

func (c *Connection) Seq() uint32 {
	return c.seq
}

func (c *Connection) Write(b []byte) (int, error) {
	n, err := c.Conn.Write(b)

	if err == nil {
		c.seq++
	}

	return n, err
}

func (c *Connection) Read(p []byte) (int, error) {
	return c.Conn.Read(p)
}
