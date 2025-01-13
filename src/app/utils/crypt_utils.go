package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSha256(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
