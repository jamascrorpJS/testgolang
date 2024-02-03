package cryptography

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"os"
)

func Validate(body []byte, digest string) (bool, error) {
	secret := os.Getenv("SECRET_KEY")
	h := hmac.New(sha1.New, []byte(secret))
	_, err := h.Write(body)
	if err != nil {
		return false, err
	}
	d := hex.EncodeToString(h.Sum(nil))
	return d == digest, nil
}
