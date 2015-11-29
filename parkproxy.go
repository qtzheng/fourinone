package fourinone

var (
	sid string
)

type ParkProxy struct {
	pk  Park
	pl  ParkLeader
	sid string
}

func NewParkProxy(host, port, parkservicecfg string) *ParkProxy {
	p := &ParkProxy{}
	p.pl = NewParkLeader(host, port, parkservicecfg, nil)
	p.pk = p.pl.getLeaderPark()
	p.init()
}
func (p *ParkProxy) init() {
	p.sid = p.pk.getSessionId()
}
func (p *ParkProxy) Create(domain string, obj Serializable) ObjectBean {
	return p.CreateWithNode(domain, GetNumber(), obj)
}
func (p *ParkProxy) CreateWithNode(domain, node string, obj Serializable) ObjectBean {
	return p.CreateWithAuth(domain, node, obj, OP_ALL)
}
func (p *ParkProxy) CreateWithAuth(domain, node string, obj Serializable, auth AuthPolicy) ObjectBean {
	return p.CreateWithAHB(domain, node, obj, auth, false)
}
func (p *ParkProxy) CreateWithHB(domain, node string, obj Serializable, heartbeat bool) ObjectBean {
	return p.CreateWithAHB(domain, node, obj, OP_ALL, heartbeat)
}
func (p *ParkProxy) CreateWithAHB(domain, node string, obj Serializable, auth AuthPolicy, heartbeat bool) ObjectBean {
	return p.CreateWithParam(domain, node, obj, auth, heartbeat, 0)
}
func (p *ParkProxy) CreateWithParam(domain, node string, obj Serializable, auth AuthPolicy, heartbeat bool, i int) ObjectBean {
	if  {
		
	}
}
func (p *ParkProxy) Update(domain, node string, obj Serializable) ObjectBean {

}
func (p *ParkProxy) Get(domain, node string) ObjectBean {

}
func (p *ParkProxy) GetLastest(domain, node string, ob ObjectBean) ObjectBean {

}
func (p *ParkProxy) GetList(domain) []ObjectBean {

}
func (p *ParkProxy) GetLastests(domain string, oblist []ObjectBean) []ObjectBean {

}
func (p *ParkProxy) Delete(domain, node string) ObjectBean {

}
func (p *ParkProxy) DeleteList(domain string) []ObjectBean {

}
func (p *ParkProxy) SetDeletable(domain string) bool {

}
func (p *ParkProxy) AddLastestListener(domain, node string, ob ObjectBean, liser LastestListener) {

}
func (p *ParkProxy) AddLastestListenerList(domain string, obList []ObjectBean, liser LastestListener) {

}
