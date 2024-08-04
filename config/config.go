package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SecretKey []byte

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

	key := os.Getenv("SECRET_KEY")
    if key == "" {
        log.Fatal("SECRET_KEY environment variable not set")
    }
    SecretKey = []byte(key)
}

func GetDsnEnv() string {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable not set")
	}

	return dsn
}
func GetSecretKey() []byte {
    return SecretKey
}

// func GetSecretKey(secretKey *[]byte) {
// 	key := os.Getenv("SECRET_KEY");
// 	if key == "" {
// 		log.Fatal("SECRET_KEY environment variable not set");
// 	}

// 	secretKey = []byte(key);

// }