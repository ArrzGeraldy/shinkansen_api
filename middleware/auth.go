package middleware

import (
	"net/http"
	"shinkansen_rest_api/config"
	"shinkansen_rest_api/exception"
	"shinkansen_rest_api/helper"
)

func ApiAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY");
		if apiKey == ""{
			panic(exception.ErrUnauthorized);
		}

		decrypt,err := helper.DecryptApiKey(apiKey);
		
		if decrypt == string(config.GetSecretKey()) && err == nil{
			next.ServeHTTP(w,r);
		} else {
			panic(exception.ErrUnauthorized);
		}

	})
}

