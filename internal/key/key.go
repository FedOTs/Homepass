package key

import (
	"golang.org/x/crypto/scrypt"
)

func GetKey(salt []byte, pass string) ([]byte, error) {
	dk, err := scrypt.Key([]byte(pass), salt, 16384, 8, 1, 32)
	return dk, err
}
