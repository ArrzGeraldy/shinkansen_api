package controller

import (
	"net/http"
	"shinkansen_rest_api/app/client/handler"
	"shinkansen_rest_api/models/web"
	"shinkansen_rest_api/service"
)

type ClientController interface {
	Index(writer http.ResponseWriter, req *http.Request);
	SignUp(writer http.ResponseWriter, req *http.Request);
	Login(writer http.ResponseWriter, req *http.Request);
	Dashboard(writer http.ResponseWriter, req *http.Request);
}

type ClientControllerImpl struct{
	AuthService service.AuthService
}

func NewClientController(service service.AuthService) ClientController{
	return &ClientControllerImpl{
		AuthService: service,
	}
}

func(controller *ClientControllerImpl) Index(writer http.ResponseWriter, req *http.Request){
	var pageData = web.PageData{}
	cookie,err := req.Cookie("session_user");

	if err == nil {
		user,err := controller.AuthService.FindUserBySession(req.Context(),cookie.Value);
		if err == nil {
			pageData = web.PageData{
				User: user,
			}
		}
		
	} 
	handler.ClientIndex(writer,pageData);
}

func(controller *ClientControllerImpl) Dashboard(writer http.ResponseWriter, req *http.Request){
	var pageData = web.PageData{}
	cookie,err := req.Cookie("session_user");
	if err != nil {
		http.Redirect(writer,req,"/",http.StatusSeeOther);
	} else {
		user,err := controller.AuthService.FindUserBySession(req.Context(),cookie.Value);
		
		if err != nil {
			http.Redirect(writer,req,"/",http.StatusSeeOther);
		} else {
			pageData = web.PageData{
				User: user,
			}

			handler.ClientDashboard(writer,pageData);
		}
	}
}

func(controller *ClientControllerImpl) SignUp(writer http.ResponseWriter, req *http.Request){

	valid := controller.checkSession(req);
	
	if !valid {
		handler.ClientSignUp(writer,nil)
		return
	} 

	http.Redirect(writer,req,"/dashboard",http.StatusSeeOther);
}

func(controller *ClientControllerImpl) Login(writer http.ResponseWriter, req *http.Request){

	valid := controller.checkSession(req);

	if !valid {
		handler.ClientLogin(writer,nil);
		return
	}

	http.Redirect(writer,req,"/dashboard",http.StatusSeeOther);
}

func (controller *ClientControllerImpl) checkSession(req *http.Request)(bool){
	cookie,err := req.Cookie("session_user");

	if err != nil {
		return false
	}

	_,err = controller.AuthService.FindUserBySession(req.Context(),cookie.Value);

	return err == nil;


}