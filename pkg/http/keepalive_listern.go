package http

import (
	"net"
	"time"
)

type KeepaliveListener struct{ net.Listener }

func (k *KeepaliveListener) Accept() (net.Conn, error) {
	c, err := k.Listener.Accept()
	if err != nil {
		return nil, err
	}
	tcpc := c.(*net.TCPConn)
	tcpc.SetKeepAlive(true)
	tcpc.SetKeepAlivePeriod(30 * time.Second)
	return tcpc, nil
}
func NewKeepaliveListener(addr string) (net.Listener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &KeepaliveListener{
		Listener: l,
	}, nil

}
