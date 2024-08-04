package middleware

import "net/http"

func CORSMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(writer http.ResponseWriter,req *http.Request){
		writer.Header().Set("Access-Control-Allow-Origin","*");
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS");

		next.ServeHTTP(writer,req);
	})
}
