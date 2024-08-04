package web

type PageData struct {
	Error string
	User  UserResponse
	Data  interface{}
}