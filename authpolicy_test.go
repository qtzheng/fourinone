package fourinone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthIncluded(t *testing.T) {
	res := AuthIncluded(1, 3)
	assert.Equal(t, res, true)
}
