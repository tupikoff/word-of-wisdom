package random

import (
	"math/rand"
	"time"
)

func FromSlice(inputs []string) string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := random.Intn(len(inputs))
	return inputs[index]
}
