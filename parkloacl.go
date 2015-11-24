package fourinone

type ParkLocal interface{
	Create(domain string,obj Serializable)ObjectBean
	CreateWithNode(domain,node string,obj Serializable)ObjectBean
	CreateWithAuth(domain,node string,obj Serializable,auth AuthPolicy)ObjectBean
	CreateWithHB(domain,node string,obj Serializable,heartbeat bool)ObjectBean
	CreateWithAHB(domain,node string,obj Serializable,auth AuthPolicy,heartbeat bool)ObjectBean
	Update(domain,node string,obj Serializable)ObjectBean
	Get(domain,node string)ObjectBean
	GetLastest(domain,node string,ob ObjectBean)ObjectBean
	GetList(domain)[]ObjectBean
	GetLastest(domain string,oblist []ObjectBean)[]ObjectBean
	Delete(domain,node string)ObjectBean
	DeleteList(domain string)[]ObjectBean
	SetDeletable(domain string)bool
	AddLastestListener(domain,node string,ob ObjectBean,liser LastestListener )
	AddLastestListenerList(domain string,obList []ObjectBean,liser LastestListener )
}