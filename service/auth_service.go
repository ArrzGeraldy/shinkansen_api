package service

import (
	"context"
	"errors"
	"shinkansen_rest_api/exception"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/models/web"
	"shinkansen_rest_api/repository"
	"time"
)

type AuthService interface {
	Register(ctx context.Context,request web.AuthRequest) (int64,error)
	Login(ctx context.Context,request web.AuthRequest,sessionKey string,expires time.Time) (int64,error)
	FindUserBySession(ctx context.Context,sessionKey string) (web.UserResponse,error);
	DestroySession(ctx context.Context,sessionKey string) (int64,error);
}

type AuthServiceImpl struct{
	Repository repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService{
	return &AuthServiceImpl{
		Repository: repo,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context,request web.AuthRequest) (int64,error)  {
	
	_,foundUser := service.Repository.FindUser(ctx,request.Username);
	if foundUser{
		return 0,exception.ErrUsernameExists
	} 

	validationValid := helper.UserValidation(request.Username,request.Password);
	if !validationValid{
		return 0,exception.ErrInvalidInput;
	}

	apiKey,err := helper.EncryptApiKey();
	helper.PanicIfError(err);

	salt,err := helper.GenerateSalt();
	helper.PanicIfError(err);

	hashedPassword := helper.HashPassword(request.Password,salt);
	return service.Repository.Register(ctx,request.Username,hashedPassword,apiKey,salt);
	
}

func(service *AuthServiceImpl) Login(ctx context.Context,request web.AuthRequest,sessionKey string,expires time.Time) (int64,error){
	
	validated := helper.UserValidation(request.Username,request.Password);
	if !validated{
		return 0,errors.New("input field min 3 and max 100");
	}
	
	user,foundUser := service.Repository.FindUser(ctx,request.Username);
	if !foundUser{
		return 0,exception.ErrLogin;
	}

	validPassword := helper.VerifyPassword(request.Password,user.Salt,user.Password);
	if !validPassword{
		return 0,exception.ErrLogin;	
	} 
	
	//return int64 and error
	return service.Repository.Login(ctx,user.Id,sessionKey,expires)
}

func(service *AuthServiceImpl) FindUserBySession(ctx context.Context,sessionKey string) (web.UserResponse,error) {
	var user web.UserResponse;

	user, err := service.Repository.FindBySession(ctx,sessionKey);
	if err != nil {
		return user,err
	}

	return user,nil;
}


func (service *AuthServiceImpl) DestroySession(ctx context.Context,sessionKey string) (int64,error){
	return service.Repository.DestroySession(ctx,sessionKey);
}



