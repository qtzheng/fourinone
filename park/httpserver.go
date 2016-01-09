package park

import (
	"fmt"
	ghttp "github.com/qtzheng/fourinone/pkg/http"
	"log"
	"net"
	"net/http"
	"time"
)

type HttpServer struct {
	router *Router
	url    string
}

func NewHttpServer(addr string, r *Router) *HttpServer {
	s := &HttpServer{}
	s.router = r
	s.url = addr
	return s
}
func (h *HttpServer) StartPeer() {
	l, err := ghttp.NewTimeoutListener(h.url, 3*time.Second, 3*time.Second)
	if err != nil {
		log.Fatalln("监听Park集群请求服务启动失败！error：", err.Error())
	}
	go func(l net.Listener) {
		fmt.Println("开始监听Peer请求！")
		log.Fatalln(h.run(l, 5*time.Minute))
	}(l)
}
func (h *HttpServer) StartClient() {
	l, err := ghttp.NewKeepaliveListener(h.url)
	if err != nil {
		log.Fatalln("监听客户端请求服务启动失败！error：", err.Error())
	}
	go func(l net.Listener) {
		fmt.Println("开始监听客户端请求！")
		log.Fatalln(h.run(l, 0))
	}(l)
}
func (h *HttpServer) run(lisn net.Listener, readTimeout time.Duration) error {
	srv := &http.Server{
		Handler:     h.router,
		ReadTimeout: readTimeout,
	}
	return srv.Serve(lisn)
}
func isMaster(c *serverContext) error {
	fmt.Println("郑乾通")
	if c.server.isMaster {
		c.response.Header().Set("Park-IsMaster", "true")
	} else {
		c.response.Header().Set("Park-IsMaster", "false")
	}
	fmt.Println("郑乾通")
	c.response.WriteHeader(http.StatusOK)
	return nil
}
