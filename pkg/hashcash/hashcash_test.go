package hashcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tupikoff/word-of-wisdom/pkg/random"
)

func TestNew(t *testing.T) {
	hc1 := New("192.168.0.1", random.String(15), 2)
	assert.True(t, hc1.IsHashValid())

	t.Log(hc1.String())

	hc2, err := NewFromString(hc1.String())
	assert.Nil(t, err)
	assert.True(t, hc2.IsHashValid())
}
