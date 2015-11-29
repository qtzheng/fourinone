package fourinone

import (
	"strconv"
)

type Server struct {
	host string
	port int
}

func NewServer(host string, port int) server {
	ser := Server{
		host: host,
		port: port,
	}
	return ser
}
func (s Server) equals(ser Server) {
	if s.host == ser.host && s.port == ser.port {
		return true
	} else {
		return false
	}
}
func (s Server) Host() string {
	return s.host
}
func (s Server) Port() int {
	return s.port
}

type ParkLeader struct {
	ismaster       bool
	alwaystry      bool
	parkservicecfg string
	thisserver     Server
	groupserver    []Server
	rpl            *AsyncExector
}

func NewParkLeader(host string, port int, parkservicecfg string, groupserver []Server) *ParkLeader {
	leader := &ParkLeader{
		ismaster:       false,
		parkservicecfg: parkservicecfg,
		thisserver:     NewServer(host, port),
	}
	alwaystry, err := strconv.ParseBool(getConfig("PARK", "ALWAYSTRYLEADER", "", "false"))
	if err != nil {
		alwaystry = false
	}
	leader.alwaystry = alwaystry
	if len(groupserver) != 0 {
		leader.groupserver = groupserver
	} else {
		leader.groupserver = [][]string{{"localhost", "1888"}, {"localhost", "1889"}, {"localhost", "1890"}}
	}
}
func (p *ParkLeader) getLeaderPark() {
	index := p.getLeaderIndex(p.thisserver)
	return p.electionLeader(-1, index)
}
func (p *ParkLeader) getNextLeader() {
	index := p.getLeaderIndex(p.thisserver)
	return p.electionLeader(index, index+1)
}
func (p *ParkLeader) getLeaderIndex(ser Server) int {
	for i := 0; i < len(p.groupserver); i++ {
		if ser.equals(p.groupserver[i]) {
			return i
		}
	}
}
func (p *ParkLeader) electionLeader(b, i int) Park {
	thesarrok := true
	if i > len(p.groupserver) {
		i = 0
	}
	server := p.groupserver[i]
	pk, ok := getBean(server.Host(), server.Port(), p.parkservicecfg).(Park)
	if ok {
		pk.askLeader()
	} else {

	}
	if thesarrok {
		p.thisserver = server
	}
	return pk
}
