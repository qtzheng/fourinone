package types

import (
	"crypto/sha1"
	"encoding/binary"
	"strconv"
)

type ID uint64

func NewID(data []byte) ID {
	sh := sha1.Sum(data)
	return ID(binary.BigEndian.Uint64(sh[:8]))
}
func (i ID) String() string {
	return strconv.FormatUint(uint64(i), 16)
}
