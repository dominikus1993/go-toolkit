package gotoolkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksumCalculationWhenObjectsAreDiferent(t *testing.T) {
	type TestObject struct {
		Name string
		Age  int
	}

	obj1 := TestObject{Name: "John", Age: 30}
	obj2 := TestObject{Name: "John", Age: 31}
	checksum1, err := CalculateChecksum(&obj1)
	assert.NoError(t, err)
	assert.NotEmpty(t, checksum1)
	checksum2, err := CalculateChecksum(&obj2)
	assert.NoError(t, err)
	assert.NotEqual(t, checksum1, checksum2)
}
