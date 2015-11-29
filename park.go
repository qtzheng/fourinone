package fourinone

type Park interface {
	create(domain, node string, bts []byte, sessionid string, auth int, heartbeat bool) ObjValue
	updateNode(domain, node string, bts []byte, sessionid string) ObjValue
	updateDomain(domain string, auth int, sessionid string) bool
	get(domain, node, sessionid string) ObjValue
	getLastest(domain, node, sessionid string, version int64) ObjValue
	delete(domain, node, sessionid string) ObjValue
	getSessionId() string
	heartbeat(domainnodekey []string, sessionid string) bool
	getParkinfo() ObjValue
	setParkinfo(ov ObjValue) bool
	askMaster() []string
	askLeader() bool
}
