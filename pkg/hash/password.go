package hash

import (
	"crypto/sha1"
	"fmt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type sha1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *sha1Hasher {
	return &sha1Hasher{salt: salt}
}

func (h *sha1Hasher) Hash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
