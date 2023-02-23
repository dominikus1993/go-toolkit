package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	str := String(10)
	assert.NotEmpty(t, str)
	assert.Len(t, str, 10)
}

func TestStringGeneration(t *testing.T) {
	str := String(10)
	str2 := String(10)
	assert.NotEqual(t, str, str2)
}
