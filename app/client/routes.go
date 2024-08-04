package client

import (
	"database/sql"
	"net/http"
	"shinkansen_rest_api/app/client/handler"
	"shinkansen_rest_api/controller"
	"shinkansen_rest_api/repository"
	"shinkansen_rest_api/service"
)



func SetupRoutes(mux *http.ServeMux, db *sql.DB) {
	authRepo := repository.NewAuthRepository(db);
	authService := service.NewAuthService(authRepo);
	authController := controller.NewAuthController(authService);
	
	clientController := controller.NewClientController(authService);

	mux.HandleFunc("POST /signup",authController.Register);
	mux.HandleFunc("POST /login",authController.Login);
	mux.HandleFunc("POST /logout",authController.Logout);

	mux.HandleFunc("/",clientController.Index);
	mux.HandleFunc("/signup",clientController.SignUp);
	mux.HandleFunc("/login",clientController.Login);
	mux.HandleFunc("/dashboard",clientController.Dashboard);
	mux.HandleFunc("/docs",func(w http.ResponseWriter, r *http.Request) {
		handler.ClientDocs(w,nil)
	});
	
	

}