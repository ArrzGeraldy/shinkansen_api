package domain

import "time"

type User struct {
	Id               	int
	Username         	string
	Password			string
	Salt               	string
	SessionKey       	string
	SessionExpiresAt 	time.Time
	ApiKey           	string
}