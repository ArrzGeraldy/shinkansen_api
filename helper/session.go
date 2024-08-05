package helper

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

func GenerateSessionKey() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CreateSession(userID int) (string, error) {
	sessionKey, err := GenerateSessionKey()
	if err != nil {
		return "", err
	}
	
	return sessionKey, nil
}

func SessionHandler(writer http.ResponseWriter,sessionKey string){
	cookie := &http.Cookie{
		Name: "session_user",
		Value: sessionKey,
		Path: "/",
		Expires:time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(writer,cookie);
}
