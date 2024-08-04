package middleware

import (
	"net/http"
	"shinkansen_rest_api/exception"
)



func ErrorMiddleware(next http.Handler) http.Handler{

	return http.HandlerFunc(func(writer http.ResponseWriter,req *http.Request){
		defer func(){
			err := recover();
			if err != nil {
				exception.ErrorHandler(writer,err);
			}
		}();
		
		next.ServeHTTP(writer,req)
	})
}

