package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CalculateSHA256(data []byte) string {
	h := sha256.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
func HashData(args ...any) string {
	var rawData string
	for _, arg := range args {
		rawData += fmt.Sprintf("%v", arg)
	}
	return CalculateSHA256([]byte(rawData))
}
