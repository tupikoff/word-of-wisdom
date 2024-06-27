package hashcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	h := hash{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	assert.True(t, h.IsValid(0))
	assert.True(t, h.IsValid(1))
	assert.True(t, h.IsValid(2))
	assert.False(t, h.IsValid(3))
	assert.False(t, h.IsValid(4))
}
