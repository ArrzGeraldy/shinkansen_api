package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"shinkansen_rest_api/config"
)

func EncryptApiKey() (string, error) {
	SecretKey := config.GetSecretKey()
	block, err := aes.NewCipher(SecretKey);
	if err != nil {
		fmt.Println(err.Error())
		panic(err);
	}

	chiperText := make([]byte, aes.BlockSize+len(SecretKey));
	iv := chiperText[:aes.BlockSize];
	_,err = io.ReadFull(rand.Reader,iv);
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block,iv);
	stream.XORKeyStream(chiperText[aes.BlockSize:],SecretKey);

	return base64.URLEncoding.EncodeToString(chiperText),nil
}

func DecryptApiKey(encryptedKey string) (string, error) {
	SecretKey := config.GetSecretKey()
	// Decode base64 encoded key
	cipherText, err := base64.URLEncoding.DecodeString(encryptedKey)
	if err != nil {
		return "", err
	}

	// Ensure the ciphertext is at least as long as the AES block size
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// Split the ciphertext into IV and encrypted message
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Create a new AES cipher block
	block, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", err
	}

	// Decrypt the key
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}



