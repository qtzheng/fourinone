package park

import (
	"net/http"
)

type serverContext struct {
	request  *http.Request
	response *Response
	server   *parkServer
	store    map[string]interface{}
}

func (s *serverContext) Request() *http.Request {
	return s.request
}
func (s *serverContext) Response() *Response {
	return s.response
}
func(s *serverContext)reset(req *http.Request, w http.ResponseWriter, srv *parkServer){
	s.request = req
	s.response.Reset(w)
	s.store = nil
	s.server = srv
}
func NewServerContext(req *http.Request, res *Response, s *parkServer) *serverContext {
	c := &serverContext{
		request:  req,
		response: res,
		server:   s,
		store:    make(map[string]interface{}),
	}
	return c
}
