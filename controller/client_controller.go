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
	var pageData = web.PageData{};

	user,valid := controller.checkSession(req);

	if valid {
		pageData.User = user;
	}

	handler.ClientIndex(writer,pageData);
}

func(controller *ClientControllerImpl) Dashboard(writer http.ResponseWriter, req *http.Request){
	var pageData = web.PageData{};

	user,valid := controller.checkSession(req);

	if valid {
		pageData.User = user;
		handler.ClientDashboard(writer,pageData);
	} else {
		http.Redirect(writer,req,"/",http.StatusSeeOther);
	}
}

func(controller *ClientControllerImpl) SignUp(writer http.ResponseWriter, req *http.Request){

	_,valid := controller.checkSession(req);
	
	if !valid {
		handler.ClientSignUp(writer,nil)
		return
	} 

	http.Redirect(writer,req,"/dashboard",http.StatusSeeOther);
}

func(controller *ClientControllerImpl) Login(writer http.ResponseWriter, req *http.Request){

	_,valid := controller.checkSession(req);

	if !valid {
		handler.ClientLogin(writer,nil);
		return
	}

	http.Redirect(writer,req,"/dashboard",http.StatusSeeOther);
}

func (controller *ClientControllerImpl) checkSession(req *http.Request) (web.UserResponse, bool) {
    cookie, err := req.Cookie("session_user")
    var userResponse web.UserResponse

    if err != nil {
        return userResponse, false
    } 
    
    userResponse, err = controller.AuthService.FindUserBySession(req.Context(), cookie.Value)
    if err != nil {
        return userResponse, false
    }
    
    return userResponse, true
}
