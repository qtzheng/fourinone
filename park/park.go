package park

const (
	configPath = "park_config.xml"
)

func StartPark() {
	c := newCluster()
	c.start()
}
