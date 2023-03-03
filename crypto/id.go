package crypto

import (
	"crypto/sha1"
	"errors"
	"strings"

	"github.com/google/uuid"
)

var errKeysAreEmpty error = errors.New("keys are empty")

func GenerateId(keys ...string) (string, error) {
	if len(keys) == 0 {
		return "", errKeysAreEmpty
	}
	key := strings.Join(keys, "-")
	data := []byte(key)
	sha1 := sha1.Sum(data)
	return uuid.NewSHA1(uuid.Nil, sha1[:]).String(), nil
}
