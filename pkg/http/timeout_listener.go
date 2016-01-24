package http

import (
	"net"
	"time"
)

type TimeoutListener struct {
	net.Listener
	wtimeoutd  time.Duration
	rdtimeoutd time.Duration
}

func (t *TimeoutListener) Accept() (net.Conn, error) {
	c, err := t.Listener.Accept()
	if err != nil {
		return nil, err
	}
	return timeoutConn{
		Conn:       c,
		wtimeoutd:  t.wtimeoutd,
		rdtimeoutd: t.rdtimeoutd,
	}, nil
}
func NewTimeoutListener(addr string, rdtimeoutd, wtimeoutd time.Duration) (net.Listener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &TimeoutListener{
		Listener:   l,
		wtimeoutd:  wtimeoutd,
		rdtimeoutd: rdtimeoutd,
	}, nil
}
