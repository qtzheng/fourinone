package fourinone

import (
	"fmt"
	"github.com/qtzheng/fourinone/park"
	"log"
	"os"
)

var (
	cfg *Config
)

const (
	cfg_path = "xmlconfig.xml"
)

func init() {
	fmt.Println("初始化配置信息")
	cfg = load_config(cfg_path)
}

func StartPark() {
	park_cfg := &park.ParkConfig{
		ParkName:        cfg.ParkName,
		Heart_beat:      cfg.Heart_beat,
		Member:          cfg.Member,
		Peers:           cfg.Peers,
		Safe_memory_per: cfg.Safe_memory_per,
	}
	var stopped <-chan struct{}
	var err error
	fmt.Println("开始执行")
	stopped, err = park.RunPark(park_cfg)
	if err != nil {
		log.Fatalln("系统出错")
	}
	<-stopped
	os.Exit(0)
}
