package main

import (
	"fmt"
	"net/http"
	"shinkansen_rest_api/app"
	"shinkansen_rest_api/app/api"
	"shinkansen_rest_api/app/client"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/middleware"
	"shinkansen_rest_api/resources"

	_ "github.com/go-sql-driver/mysql"
)



func main() {
	db := app.NewDB();

	var mux *http.ServeMux = http.NewServeMux();

	resources.FileServer(mux);

	client.SetupRoutes(mux,db);
	api.SetupRoutes(mux,db);

	handler := middleware.CORSMiddleware(middleware.ErrorMiddleware(mux));

	fmt.Println("Server run on PORT 4000");

	server := http.Server{
		Addr: "localhost:4000",
		Handler: handler,
	}

	err := server.ListenAndServe();

	helper.PanicIfError(err);
}