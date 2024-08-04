package helper

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

// GenerateSalt menghasilkan salt acak
func GenerateSalt() (string, error) {
	salt := make([]byte, 16) // 16 byte salt
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword menghasilkan hash dari password dengan salt
func HashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// VerifyPassword memverifikasi password dengan hash yang diberikan dan salt
func VerifyPassword(password, salt, hashedPassword string) bool {
	return HashPassword(password, salt) == hashedPassword
}
