package park

import (
	"time"
)

type ParkConfig struct {
	//park_config
	ParkName        string
	Member          *Member
	Peers           map[string]*Member
	Heart_beat      time.Duration
	Safe_memory_per int
}
