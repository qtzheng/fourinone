package fourinone

import (
	"github.com/satori/go.uuid"
	"os"
	"runtime"
	"strings"
)

func SetConfigFile(cfgFile string) {
	configFile = cfgFile
}
func StartPark() {
	parkcfg := getParkConfig()
}
func StartParkWithCfg(host string, port int, sn string, servers [][]string) {
	if sn == "" {
		sn = getParkService()
	}
	if getConfig("PARK", "STARTWEBAPP", "", "false") == "true" {
		startInetServer()
	}
}
func startInetServer() {
	inetcfg := getInetConfig()
}
func GetPark() ParkLocal {
	parkCfg := getParkConfig()
}
func GetNumber() string {
	uuid := uuid.NewV1().String()
	uuid = strings.Replace(uuid, "-", "", -1)
	return uuid
}
func GetWindows(defaul bool) bool {
	if !defaul {
		return runtime.GOOS == "windows"
	}
	return defaul
}
func GetLinux() bool {
	return runtime.GOOS == "linux"
}
