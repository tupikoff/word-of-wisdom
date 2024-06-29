package hashcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	h := hash{1, 0, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	assert.True(t, h.IsValid(0))
	assert.True(t, h.IsValid(7))
	assert.False(t, h.IsValid(8))

	h = hash{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	assert.True(t, h.IsValid(0))
	assert.True(t, h.IsValid(23))
	assert.False(t, h.IsValid(24))

	h = hash{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	assert.True(t, h.IsValid(0))
	assert.True(t, h.IsValid(23))
	assert.False(t, h.IsValid(24))
}
