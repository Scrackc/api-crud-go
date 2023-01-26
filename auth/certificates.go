package auth

import (
	"crypto/rsa"
	"os"
	"sync"

	"github.com/golang-jwt/jwt/v4"
)

var (
	singKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	onse      sync.Once
)

// LoadFiles .
func LoadFiles(privateFilePath, publicFilePath string) error {
	var err error
	onse.Do(func() {
		err = loadFiles(privateFilePath, publicFilePath)
	})
	return err
}

func loadFiles(privateFilePath, publicFilePath string) error {
	privateBytes, err := os.ReadFile(privateFilePath)
	if err != nil {
		return err
	}

	publicBytes, err := os.ReadFile(publicFilePath)
	if err != nil {
		return err
	}
	return parseRSA(privateBytes, publicBytes)
}

func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	singKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}
	return nil
}
