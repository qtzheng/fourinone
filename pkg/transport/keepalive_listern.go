package transport

import (
	"net"
)

type keepaliveListener struct{ net.Listener }
