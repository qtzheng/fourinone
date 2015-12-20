package park

import (
	"encoding/xml"
	"github.com/qtzheng/fourinone/pkg/types"
	gxml "github.com/qtzheng/fourinone/pkg/xml"
	"github.com/qtzheng/fourinone/store"
	"time"
)

const (
	configPath = "config.xml"
)

type cluster struct {
	id    types.ID
	name  string
	store store.Store
	sync.Mutex
	member          *member
	peers           map[types.ID]*member
	version         uint64
	heart_beat      time.Duration
	safe_memory_per int
}

func (c *cluster) Init(token xml.Token) {
	switch t := token.(type) {
	case xml.StartElement:
		switch t.Name.Local {
		case "ParkCluster":
			for _, attr := range t.Attr {
				switch attr.Name.Local {
				case "Name":
					c.name = attr.Value
				case "HeartBeat":
					c.heart_beat = attr.Value * time.Second
				case "SafeMemoryPer":
					c.safe_memory_per = attr.Value
				}
			}
		case "Member", "Peer":
			var (
				name, clientUrl, peerUrl string
			)
			for _, attr := range t.Attr {
				switch attr.Name.Local {
				case "Name":
					name = attr.Value
				case "ClientUrl":
					clientUrl = attr.Value
				case "PeerUrl":
					peerUrl = attr.Value
				}
			}
			m := newMember(name, peerUrl, clientUrl, time.Now())
			if t.Name.Local == "Member" {
				c.member = m
			} else {
				c.peers[m.ID] = m
			}
		case "":
		}
	case xml.EndElement:
	}
}
func (c *cluster) start() {
	c.member
}
func newCluster() *cluster {
	c := &cluster{
		peers: make(map[types.ID]*member),
		store: store.New(),
	}
	gxml.ParseXmlConfig(configPath, c)
}
