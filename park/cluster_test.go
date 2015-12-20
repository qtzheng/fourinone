package park

import (
	"fmt"
	"testing"
)

func Test_newCluster(t *testing.T) {
	c := newCluster()
	fmt.Println(c)
}
