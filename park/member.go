package park

import (
//"fmt"
//"time"
)

type Member struct {
	Name      string `json:"name"`
	PeerUrl   string `json:"peerUrl"`
	ClientUrl string `json:"clientUrl,omitempty"`
}

func NewMember(name, peerUrl, clientUrl string) *Member {
	m := &Member{
		Name:      name,
		PeerUrl:   peerUrl,
		ClientUrl: clientUrl,
	}
	return m
}
func (m *Member) equal(memb *Member) bool {
	if memb == nil {
		return false
	}
	return m.Name == memb.Name
}
