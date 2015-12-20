package transport

import (
	"net"
	"time"
)

type timeoutListener struct {
	net.Listener
	wtimeoutd  time.Duration
	rdtimeoutd time.Duration
}

func (t *timeoutListener) Accept() (net.Conn, error) {
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
	l, err := NewListener(addr)
	if err != nil {
		return nil, err
	}
	return &timeoutListener{
		Listener:   l,
		wtimeoutd:  wtimeoutd,
		rdtimeoutd: rdtimeoutd,
	}, nil
}
