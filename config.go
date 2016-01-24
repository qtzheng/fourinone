package fourinone

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/qtzheng/fourinone/park"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Config struct {
	//park_config
	ParkName        string
	Member          *park.Member
	Peers           map[string]*park.Member
	Heart_beat      time.Duration
	Safe_memory_per int
}

func load_config(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("开始生成配置文件！")
			data = []byte(getDefaultConfig())
			f, f_err := os.Create(path)
			defer f.Close()
			if f_err != nil {
				fmt.Println("创建配置文件生成失败,继续运行！")
			}
			if _, err = f.Write(data); err != nil {
				fmt.Println("配置文件写入默认配置失败，继续运行！")
			} else {
				fmt.Println("配置文件生成成功！")
			}
		} else {
			fmt.Printf("读取配置文件:%s,失败:%s\n", path, err.Error())
			os.Exit(0)
		}
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(data))
	c := &Config{
		Peers: make(map[string]*park.Member),
	}
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		parseXmlToken(t, c)
	}
	return c
}
func parseXmlToken(token xml.Token, c *Config) {
	switch t := token.(type) {
	case xml.StartElement:
		switch t.Name.Local {
		case "ParkCluster":
			for _, attr := range t.Attr {
				switch attr.Name.Local {
				case "Name":
					c.ParkName = attr.Value
				case "HeartBeat":
					v, err := strconv.Atoi(attr.Value)
					if err != nil {
						v = 3
					}
					c.Heart_beat = time.Duration(v) * time.Second
				case "SafeMemoryPer":
					v, err := strconv.Atoi(attr.Value)
					if err != nil {
						v = 90
					}
					c.Safe_memory_per = v
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
			m := park.NewMember(name, peerUrl, clientUrl)
			if t.Name.Local == "Member" {
				c.Member = m
			} else {
				c.Peers[m.Name] = m
			}
		case "":
		}
	case xml.EndElement:
	}
}
func getDefaultConfig() string {
	config := `<?xml version="1.0" encoding="utf-8" ?>
<Config>
  <ParkCluster HeartBeat="3" MaxDelay="3" SafeMemoryPer="95">
    <Member Name="park1" ClientUrl="localhost:1500" PeerUrl="localhost:1501"></Member>
    <Peer Name="park2" ClientUrl="localhost:2000" PeerUrl="localhost:2001"></Peer>
    <Peer Name="park3" ClientUrl="localhost:3000" PeerUrl="localhost:3001"></Peer>
    <Peer Name="park4" ClientUrl="localhost:4000" PeerUrl="localhost:4001"></Peer>
  </ParkCluster>
</Config>`
	return config
}
