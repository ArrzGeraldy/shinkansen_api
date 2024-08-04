package middleware

import (
	"encoding/json"
	"net/http"
	"shinkansen_rest_api/config"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/models/web"
)

func ApiAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY");
		if apiKey == ""{
			w.WriteHeader(http.StatusUnauthorized);
			response := web.WebResponse{
				Code: 401,
				Status: "Unauthorized",
				Error: "Unauthorized",
			}

			encoder := json.NewEncoder(w);
			err := encoder.Encode(response);
			helper.PanicIfError(err);
			return
		}

		decrypt,err := helper.DecryptApiKey(apiKey);
		
		if decrypt == string(config.GetSecretKey()) && err == nil{
			next.ServeHTTP(w,r);
		} else {
			w.WriteHeader(http.StatusUnauthorized);
			response := web.WebResponse{
				Code: 401,
				Status: "Unauthorized",
				Error: "Invalid key",
			}

			encoder := json.NewEncoder(w);
			err := encoder.Encode(response);
			helper.PanicIfError(err);
			return
		}

	})
}

