package exception

import (
	"errors"
	"net/http"
	"shinkansen_rest_api/app/client/handler"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/models/web"
)

func ErrorHandler(writer http.ResponseWriter, err interface{}) {
	if exception, ok := err.(error); ok {
		switch {
			case errors.Is(exception, ErrNotFound):
				notFoundError(writer, exception)

			case errors.Is(exception, ErrParameter):
				parameterError(writer, exception)

			case errors.Is(exception, ErrLogin):
				loginInvalid(writer,exception);

			case errors.Is(exception, ErrUnauthorized):
				unauthorizedErr(writer,exception);

			default:
				internalServerError(writer, exception)
		}
	}
}

func notFoundError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusNotFound)

	response := web.WebResponse{
		Code:   http.StatusNotFound,
		Status: "NOT FOUND",
		Error:  err.Error(),
	}

	helper.ToResponseJson(writer,response)
}


func loginInvalid(writer http.ResponseWriter,err error){
	data := web.PageData{
		Error: err.Error(),
	}
	handler.ClientLogin(writer,data)
}

func parameterError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusBadRequest)

	response := web.WebResponse{
		Code:   http.StatusBadRequest,
		Status: "BAD REQUEST",
		Error:  err.Error(),
	}

	helper.ToResponseJson(writer,response);
}

func unauthorizedErr(writer http.ResponseWriter, err error){
	writer.WriteHeader(http.StatusUnauthorized);
	response := web.WebResponse{
		Code: 401,
		Status: "Unauthorized",
		Error: err.Error(),
	};

	helper.ToResponseJson(writer,response);
}

func internalServerError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusInternalServerError)

	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Error:  err.Error(),
	}

	helper.ToResponseJson(writer,response);
}
