package controller

import (
	"net/http"
	"shinkansen_rest_api/app/client/handler"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/models/web"
	"shinkansen_rest_api/service"
	"time"
)

type AuthController interface {
	Register(writer http.ResponseWriter,req *http.Request);
	Login(writer http.ResponseWriter,req *http.Request);
	Logout(writer http.ResponseWriter,req *http.Request);
}

type AuthControllerImpl struct{
	Service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController{
	return &AuthControllerImpl{
		Service: service ,
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter,req *http.Request){
	reqForm := web.AuthRequest{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}
	
	isValid,err := controller.Service.Register(req.Context(),reqForm);
	if err != nil {
		handler.ClientSignUp(writer,web.PageData{
			Error: err.Error(),
		});
		return;
	}
	
	if isValid > 0 {
		controller.Login(writer,req);
	} 
}
func (controller *AuthControllerImpl) Login(writer http.ResponseWriter,req *http.Request){
	reqForm := web.AuthRequest{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	sessionKey,err := helper.GenerateSessionKey();
	helper.PanicIfError(err);

	expiresSession := time.Now().Add(224 * time.Hour);
	isValid,err := controller.Service.Login(req.Context(),reqForm,sessionKey,expiresSession);

	if err != nil {
		handler.ClientLogin(writer,web.PageData{
			Error: err.Error(),
		});
		return;
	}

	if isValid > 0 {
		helper.SessionHandler(writer,sessionKey);
		http.Redirect(writer,req,"/dashboard",http.StatusSeeOther);
	}
}

func(controller *AuthControllerImpl) Logout(writer http.ResponseWriter,req *http.Request){

	cookie,err := req.Cookie("session_user");
	if err == nil {
		_,err := controller.Service.DestroySession(req.Context(),cookie.Value);
		if err == nil {
			http.SetCookie(writer, &http.Cookie{
                Name:   "session_user",
                Value:  "",
                Path:   "/", // Adjust path if necessary
                Expires: time.Unix(0, 0), // Set expiration date to the past
            })
			http.Redirect(writer,req,"/",http.StatusSeeOther);
		}
	}

}

