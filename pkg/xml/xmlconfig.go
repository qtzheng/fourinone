package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type XmlAnalysis interface {
	Init(t xml.Token)
}

func ParseXmlConfig(path string, a XmlAnalysis) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("开始生成配置文件！")
			data = []byte(getDefaultConfig())
			f, f_err := os.Create(path)
			defer f.Close()
			if f_err != nil {
				fmt.Println("配置文件生成失败！")
			}
			if _, err = f.Write(data); err != nil {
				fmt.Println("配置文件生成失败！")
			} else {
				fmt.Println("配置文件生成成功！")
			}
		} else {
			fmt.Printf("读取配置文件:%s,失败:%s\n", path, err.Error())
			os.Exit(0)
		}

	}
	parseXmlConfig(data, a)
}
func parseXmlConfig(data []byte, a XmlAnalysis) {
	decoder := xml.NewDecoder(bytes.NewBuffer(data))
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		a.Init(t)
	}
}
func getDefaultConfig() string {
	config := `<?xml version="1.0" encoding="utf-8" ?>
<Config>
  <ParkCluster HeartBeat="3" MaxDelay="3" SafeMemoryPer="95">
    <Member Name="park1" ClientUrl="localhost:1000" PeerUrl="localhost:1001"></Member>
    <Member Name="park2" ClientUrl="localhost:2000" PeerUrl="localhost:2001"></Member>
    <Member Name="park3" ClientUrl="localhost:3000" PeerUrl="localhost:3001"></Member>
    <Member Name="park4" ClientUrl="localhost:4000" PeerUrl="localhost:4001"></Member>
  </ParkCluster>
</Config>`
	return config
}
