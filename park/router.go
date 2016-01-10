package park

import (
	"net/http"
	"strings"
)

const (
	ContentType                = "Content-Type"
	TextPlain                  = "text/plain"
	ApplicationJSON            = "application/json"
	CharsetUTF8                = "charset=utf-8"
	ApplicationJSONCharsetUTF8 = ApplicationJSON + "; " + CharsetUTF8

	DELETE = "DELETE"
	// GET HTTP method
	GET = "GET"
	// HEAD HTTP method
	HEAD = "HEAD"
	// OPTIONS HTTP method
	OPTIONS = "OPTIONS"
	// PATCH HTTP method
	PATCH = "PATCH"
	// POST HTTP method
	POST = "POST"
	// PUT HTTP method
	PUT = "PUT"
	// TRACE HTTP method
	TRACE = "TRACE"
)

type HandlerFunc func(c *serverContext) error
type Route struct {
	delete  HandlerFunc
	get     HandlerFunc
	head    HandlerFunc
	options HandlerFunc
	patch   HandlerFunc
	post    HandlerFunc
	put     HandlerFunc
	trace   HandlerFunc
}

func (r *Route) find(method string) HandlerFunc {
	switch method {
	case GET:
		return r.get
	case DELETE:
		return r.delete
	case HEAD:
		return r.head
	case OPTIONS:
		return r.options
	case PATCH:
		return r.patch
	case POST:
		return r.post
	case PUT:
		return r.put
	case TRACE:
		return r.trace
	default:
		return nil
	}
}

type Router struct {
	prefix string //前缀
	routes map[string]*Route
	server *parkServer
}

func NewRouter(prefix string,s *parkServer) *Router {
	r := &Router{
		prefix: prefix,
		routes: make(map[string]*Route),
		server:s,
	}
	return r
}
func (r *Router) find(method, path string) (HandlerFunc, bool) {
	path = strings.TrimLeft(path, r.prefix)
	route, ok := r.routes[path]
	if !ok {
		return nil, false
	}
	h := route.find(method)
	if h == nil {
		return nil, false
	}
	return h, true
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h, ok := r.find(req.Method, req.URL.Path)
	if !ok {
		w.Header().Set(ContentType, TextPlain)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("无效的请求"))
		return
	}
	c := r.server.pool.Get().(*serverContext)
	c.reset(req,w,r.server)
	if err := h(c); err != nil {
		w.Header().Set(ContentType, TextPlain)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("请求执行失败"))
	}
}
func (r *Router) Add(action, path string, h HandlerFunc) {
	v, ok := r.routes[path]
	if !ok {
		v = new(Route)
	}
	switch action {
	case GET:
		v.get = h
	case POST:
		v.post = h
	case DELETE:
		v.delete = h
	}
	r.routes[path]=v
}
