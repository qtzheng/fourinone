package park

import (
	"fmt"
	"github.com/qtzheng/fourinone/pkg/types"
	"time"
)

type member struct {
	ID        types.ID `json:"id"`
	Name      string   `json:"name,omitempty"`
	PeerUrl   string   `json:"peerUrl"`
	ClientUrl string   `json:"clientUrl,omitempty"`
}

func newMember(name, peerUrl, clientUrl string, now time.Time) *member {
	m := &member{
		Name:      name,
		PeerUrl:   peerUrl,
		ClientUrl: clientUrl,
	}
	data := []byte(name)
	data = append(data, []byte(peerUrl)...)
	data = append(data, []byte(clientUrl)...)
	if now != nil {
		data = append(data, []byte(fmt.Sprintf("%d", now.Unix()))...)
	}
	m.ID = types.NewID(data)
}
func (m *member) name() {

}
