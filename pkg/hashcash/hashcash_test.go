package hashcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	hc1 := New("192.168.0.1")
	assert.True(t, hc1.Hash().IsValid())

	hc2, err := NewFromString(hc1.String())
	assert.Nil(t, err)
	assert.True(t, hc2.Hash().IsValid())
}
