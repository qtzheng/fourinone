package http

import (
	"net"
	"time"
)

type timeoutConn struct {
	net.Conn
	wtimeoutd  time.Duration
	rdtimeoutd time.Duration
}

func (t timeoutConn) Write(b []byte) (n int, err error) {
	if t.wtimeoutd > 0 {
		if err := t.SetWriteDeadline(time.Now().Add(t.wtimeoutd)); err != nil {
			return 0, err
		}
	}
	return t.Conn.Write(b)
}
func (t timeoutConn) Read(b []byte) (n int, err error) {
	if t.rdtimeoutd > 0 {
		if err := t.SetReadDeadline(time.Now().Add(t.rdtimeoutd)); err != nil {
			return 0, err
		}
	}
	return t.Conn.Read(b)
}
