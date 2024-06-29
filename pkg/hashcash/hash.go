package hashcash

import (
	"encoding/hex"
)

type hash [20]byte

func (h hash) IsValid(zeroBitNum int) bool {
	var zeroCount int
	for _, hashValue := range h {
		for bitIndex := 7; bitIndex >= 0; bitIndex-- {
			if zeroCount == zeroBitNum {
				return true
			}
			bit := hashValue & (1 << bitIndex)
			if bit != 0 {
				return false
			}
			zeroCount++
		}
	}
	if zeroCount == zeroBitNum {
		return true
	}
	return false
}

func (h hash) String() string {
	return hex.EncodeToString(h[:])
}
