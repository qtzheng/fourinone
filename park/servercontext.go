package park

import (
	"net/http"
	"net/url"
)

type serverContext struct {
	request  *http.Request
	response *Response
	server   *parkServer
	store    map[string]interface{}
	query    url.Values
}

func (s *serverContext) Request() *http.Request {
	return s.request
}
func (s *serverContext) Response() *Response {
	return s.response
}
func (s *serverContext) reset(req *http.Request, w http.ResponseWriter, srv *parkServer) {
	s.request = req
	s.response.Reset(w)
	s.store = nil
	s.server = srv
	s.query = nil
}
func (s *serverContext) Query(name string) string {
	if s.query == nil {
		s.query = s.request.URL.Query()
	}
	return s.query.Get(name)
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
