package gotest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type AB struct {
	a int
	B int
}

// add a log
func countProof(token, secret string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(token))
	if err != nil {
		return "", err
	}
	expectedMAC := mac.Sum(nil)
	return hex.EncodeToString(expectedMAC), nil
}
