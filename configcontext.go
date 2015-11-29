package fourinone

import (
	"strings"
)

var (
	mb                                                                                     MulBean
	c_QSXYSJ, c_YMMZ, c_RZDY, c_YCDYXY, c_DMY, c_AQCL, c_POLICY, c_LSML, c_SERVICEONWORKER string
	c_TMOT                                                                                 int64 = -1
	c_USERS                                                                                ObjValue
	c_KEYLENTH                                                                             int64 = -1
	c_VALUELENGTH                                                                          int64 = -1
	c_REGIONLENGTH                                                                         int64 = -1
	c_LOADLENGTH                                                                           int64 = -1
	c_HASHCAPACITY                                                                         int   = -1
	c_DATAROOT                                                                             string
)

const configFile = "config.xml"

func getKEYLENTH() int64 {

}

func getVALUELENGTH() int64 {
	return 0
}
func getREGIONLENGTH() int64 {
	return 0
}
func getLOADLENGTH() int64 {
	return 0
}
func getHASHCAPACITY() int {
	return 0
}
func getDATAROOT() string {
	return ""
}
func getMulBean() MulBean {
	return nil
}
func getQSXYSJ() string {
	return ""
}
func getYMMZ() string {
	return ""
}
func getRZDY() string {
	return ""
}
func getYCDYXY() string {
	return ""
}
func getDMY() string {
	return ""
}
func getAQCL() string {
	return ""
}
func getPOLICY() string {
	return ""
}
func getLSML() string {
	return ""
}
func getProp(propstr string) string {
	return ""
}
func getProtocolInfo(ym string, dk int, mc string) string {
	return ""
}
func getTMOT() int64 {
	return 0
}
func getServiceFlag() bool {
	return false
}
func getSecTime(hours float32) int64 {
	return int64(hours * 3600 * 1000)
}
func getParkConfig() [][]string {
	servers := getConfig("PARK", "SERVERS", "", "")
	return getServerFromStr(servers)
}
func getParkService() string {
	return getConfig("PARK","SERVICE","","")
}
func getCtorService() []string {
	return nil
}
func getFttpConfig() []string {
	return nil
}

func getInetConfig() []string {
	return strings.Split(getConfig("WEBAPP","SERVERS","",""),":")
}
func getUsersConfig() ObjValue {
	return nil
}
func getInetStrConfig() string {
	return ""
}
func getPolicyConfig() string {
	return ""
}
func getWorkerConfig() []string {
	return nil
}
func getCacheConfig() [][]string {
	return nil
}
func getCacheService() string {
	return ""
}
func getCacheFacadeConfig() []string {
	return nil
}
func getCacheFacadeService() string {
	return ""
}
func getInitServices() int {
	return 0
}
func getMaxServices() int {
	return 100
}
func getParallelPattern() int {
	return 0
}
func getConfig(cfgname, cfgprop, cfgdesc, defvalue string) string {
	xu := &XmlUtil{}
	al := xu.GetXmlObjectByFile(configFile, cfgname, cfgdesc)

	if al != nil && len(al) > 0 {
		cfgProps := al[0]
		cfgProps.GetData(cfgprop, defvalue)
	}
	return ""
}
func getLogLevel(deflevel string) string {
	return ""
}
func getCacheGroupConfig() ObjValue {
	return nil
}
func getDateLong(dateStr string) string {
	return ""
}
func getServerFromStr(servers string) [][]string {
	serverarr := strings.Split(servers, ",")
	sarr := make([][]string, len(serverarr))
	for n, v := range serverarr {
		hostport := strings.Split(serverarr[n], ":")
	}
	return sarr
}
func getObjFromStr(strs string) ObjValue {
	serverarr := strings.Split()
}
func getRequest(requestUrl string) string {
	return ""
}
