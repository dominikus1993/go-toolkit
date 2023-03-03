package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	subject, err := GenerateId("jan", "pawel", "II")
	assert.NoError(t, err)
	assert.Equal(t, "29cb24d3-4364-5479-8cf0-98fed5fb6141", subject)
}

func TestUUIDReapetable(t *testing.T) {
	keys := []string{"jan", "pawel", "II"}
	subject, err := GenerateId(keys...)
	assert.NoError(t, err)
	subject2, err := GenerateId(keys...)
	assert.NoError(t, err)
	assert.Equal(t, subject2, subject)
}

func TestUUIDWhenKeysAreEmpty(t *testing.T) {
	_, err := GenerateId()
	assert.Error(t, err)
}

func BenchmarkGenerateId(b *testing.B) {
	keys := []string{"jan", "pawel", "II"}
	for i := 0; i < b.N; i++ {
		GenerateId(keys...)
	}
}
