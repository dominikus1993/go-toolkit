package gotoolkit

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
)

func CalculateChecksum(body any) (string, error) {
	var b bytes.Buffer
	err := gob.NewDecoder(&b).Decode(body)
	if err != nil {
		return "", fmt.Errorf("error calculating checksum: %w", err)
	}
	result := md5.Sum(b.Bytes())
	return fmt.Sprintf("%x", result), nil
}
