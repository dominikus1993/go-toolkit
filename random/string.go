package random

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) (string, error) {
	if length < 0 {
		return "", errors.New("length smaller than 0")
	}
	b := make([]byte, length)
	len := big.NewInt(int64(len(charset)))
	for i := range b {
		a, err := rand.Int(rand.Reader, len)
		if err != nil {
			return "", err
		}
		b[i] = charset[a.Int64()]
	}
	return string(b), nil
}

func String(length int) (string, error) {
	return StringWithCharset(length, charset)
}
