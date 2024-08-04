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
	// 2024-08-04T09:45:28.627Z
}

// func validateSession(sessionKey string) (int, error) {
// 	var userID int
// 	var expiresAt time.Time

// 	// Periksa session_key di database
// 	err := db.QueryRow("SELECT user_id, expires_at FROM sessions WHERE session_key = ?", sessionKey).Scan(&userID, &expiresAt)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Periksa apakah sesi sudah kadaluwarsa
// 	if time.Now().After(expiresAt) {
// 		return 0, fmt.Errorf("session expired")
// 	}

// 	return userID, nil
// }

// func sessionHandler(w http.ResponseWriter, r *http.Request) {
// 	sessionKey := r.Header.Get("Session-Key")
// 	userID, err := validateSession(sessionKey)
// 	if err != nil {
// 		http.Error(w, "Invalid session", http.StatusUnauthorized)
// 		return
// 	}

// 	fmt.Fprintf(w, "Welcome User ID: %d", userID)
// }
