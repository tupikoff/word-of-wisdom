package hashcash

import (
	"encoding/hex"
)

type hash [20]byte

func (h hash) IsValid(zeroBitNum int) bool {
	for i := range zeroBitNum {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

func (h hash) String() string {
	return hex.EncodeToString(h[:])
}
