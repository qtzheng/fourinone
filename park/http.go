package park

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

const (
	membersHandlerPath = "/members"
)

func serveHTTP(l net.Listener, handler http.Handler, readTimeout time.Duration) error {
	srv := &http.Server{
		Handler:     handler,
		ReadTimeout: readTimeout,
	}
	return srv.Serve(l)
}

type membersHandler struct {
	cluster Cluster
}

func (m *membersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("F-Cluster-ID", m.cluster.ID().String())
	if r.URL.Path != membersHandlerPath {
		http.Error(w, "bad path", http.StatusBadRequest)
		return
	}
	ms := m.cluster.Members()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ms); err != nil {
		fmt.Printf("failed to encode members response (%v)", err)
	}
}
func NewPeerHandler(cluster Cluster) http.Handler {
	mh := &membersHandler{
		cluster: cluster,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.NotFound)
	mux.Handle(membersHandlerPath, mh)
	return mux
}
