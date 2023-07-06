package generator

import (
	"crypto/rand"
	"math/big"
)

const (
	characterSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%&*()?{}[]"
)

func GeneratePassword(length int) string {
	charsetLength := big.NewInt(int64(len(characterSet)))
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, charsetLength)
		password[i] = characterSet[randomIndex.Int64()]
	}

	return string(password)
}
