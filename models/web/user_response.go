package web

import "time"

type UserResponse struct {
	Id               int
	Username         string
	SessionKey       string
	SessionExpiresAt time.Time
	ApiKey           string
}