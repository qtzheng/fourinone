package park

import (
	"fmt"
	ghttp "github.com/qtzheng/fourinone/pkg/http"
	"net/http"
	"sync"
	"time"
	"github.com/qtzheng/fourinone/store"
)

const (
	prefix       = ""
	isMasterPath = "/ismaster"
)

type Park interface {
	Run() (<-chan struct{}, error)
	AskMaster()
	//GetLastest()
}
type parkServer struct {
	cfg      *ParkConfig
	isMaster bool
	mu       sync.RWMutex
	pool     sync.Pool
	stop     chan struct{}
	store    store.Store
}

func newParkServer(cfg *ParkConfig) *parkServer {
	s := &parkServer{
		cfg:      cfg,
		isMaster: false,
	}
	s.pool.New = func() interface{} {
		return NewServerContext(nil, new(Response), s)
	}
	s.store=store.New()
	return s
}
func RunPark(cfg *ParkConfig) (<-chan struct{}, error) {
	s := newParkServer(cfg)
	p_router := NewRouter(prefix,s) //用于集群间通信路由
	p_router.Add(GET, isMasterPath, isMaster)
	c_router := NewRouter(prefix,s) //用于客户端通信路由
	p_httpServer := NewHttpServer(cfg.Member.PeerUrl, p_router)
	p_httpServer.StartPeer()

	c_httpServer := NewHttpServer(cfg.Member.ClientUrl, c_router)
	c_httpServer.StartClient()
	return s.Run()
}
func (p *parkServer) Run() (<-chan struct{}, error) {
	p.wantBeMaster()
	return nil, nil
}
func (p *parkServer) wantBeMaster() {
	fmt.Printf("%s-->集群：%s-->节点：%s-->请求成为主节点!\n", time.Now().Format("2006-01-02 15:04:05"), p.cfg.ParkName, p.cfg.Member.Name)
	var (
		client ghttp.HttpClient
		err    error
		resp   *http.Response
		rest   string
		master *Member
	)
	for _, peer := range p.cfg.Peers {
		fmt.Printf("%s-->集群：%s-->节点：%s-->询问节点【%s】是否为主节点!\n", time.Now().Format("2006-01-02 15:04:05"), p.cfg.ParkName, p.cfg.Member.Name, peer.Name)
		client = ghttp.NewPeerClient(GET, isMasterPath, peer.PeerUrl)
		resp, err = client.Do()
		if err != nil {
			fmt.Printf("%s-->询问失败，原因：%s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
			continue
		} else {
			fmt.Printf("%s-->询问成功，：%s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
			rest = resp.Header.Get(isMasterPath)
			if rest == "true" {
				master = peer
			}
		}
	}
	if master == nil {
		p.isMaster = true
	}
}
func (p *parkServer) AskMaster() {

}
