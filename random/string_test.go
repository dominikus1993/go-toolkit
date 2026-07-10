package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	str, err := String(10)
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
	assert.Len(t, str, 10)
}

func TestStringGeneration(t *testing.T) {
	str, err := String(10)
	assert.NoError(t, err)
	str2, err := String(10)
	assert.NoError(t, err)
	assert.NotEqual(t, str, str2)
}
