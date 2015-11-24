package fourinone

func SetConfigFile(cfgFile string) {
	configFile = cfgFile
}
func GetPark() ParkLocal {
	parkCfg := getParkConfig()
}
