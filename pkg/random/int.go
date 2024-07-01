package random

import (
	"math/rand"
	"time"
)

func IntIn(min, max int) int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + random.Intn(max-min+1)
}

func Int() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Int()
}
