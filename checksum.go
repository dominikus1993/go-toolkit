package gotoolkit

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
)

func CalculateChecksum[T any](body *T) (string, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(body)
	if err != nil {
		return "", fmt.Errorf("error calculating checksum: %w", err)
	}
	result := md5.Sum(b.Bytes())
	return fmt.Sprintf("%x", result), nil
}
