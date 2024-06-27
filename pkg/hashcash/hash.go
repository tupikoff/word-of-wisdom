package hashcash

import "encoding/hex"

type hash [20]byte

func (h hash) IsValid() bool {
	if h[0] == 0 &&
		h[1] == 0 &&
		h[2] <= 15 {
		return true
	}
	return false
}

func (h hash) String() string {
	return hex.EncodeToString(h[:])
}
