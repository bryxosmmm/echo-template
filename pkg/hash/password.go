package hash

import (
	"errors"

	"github.com/matthewhartstonge/argon2"
)

func GenerateHash(password string) (string, error) {
	config := argon2.DefaultConfig()
	hash, err := config.HashEncoded([]byte(password))
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(password, hash string) error {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hash))
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("passwords do not match")
	}

	return nil
}
