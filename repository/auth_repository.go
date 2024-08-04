package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shinkansen_rest_api/exception"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/models/domain"
	"shinkansen_rest_api/models/web"
	"time"
)

type AuthRepository interface {
	Register(ctx context.Context,username string, password string,apiKey string,salt string) (int64,error);
	FindUser(ctx context.Context,username string) (domain.User,bool);
	Login(ctx context.Context,id int,sessionKey string, expires time.Time) (int64,error);
	FindBySession(ctx context.Context,sessionKey string) (web.UserResponse,error);
	DestroySession(ctx context.Context,sessionKey string) (int64,error);
}

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository{
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (repo *AuthRepositoryImpl) Register(ctx context.Context, username string, password string,apiKey string,salt string) (int64,error) {
	stmt := "INSERT INTO users (username, password,salt,api_key) VALUES(?, ?, ?, ?)"
	rows, err := repo.DB.ExecContext(ctx, stmt, username, password,salt,apiKey)

	if err != nil {
		fmt.Println("Error executing query:", err.Error())
		return 0,exception.ErrServer
	}

	row, err := rows.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err.Error())
		return 0,exception.ErrServer
	}

	return row,nil
}

func (repo *AuthRepositoryImpl) Login(ctx context.Context, id int,sessionKey string, expires time.Time) (int64,error) {
	stmt := "UPDATE users SET session_key = ?, session_expires_at = ? WHERE id = ?"
	result, err := repo.DB.ExecContext(ctx, stmt, sessionKey,expires,id);
	if err != nil {
		return 0,exception.ErrServer
	}

	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err.Error())
		return 0,errors.New("login server error");
	}

	return row,nil

}

func (repo *AuthRepositoryImpl) FindUser(ctx context.Context, username string) (domain.User,bool) {
	stmt := "SELECT id,username,password,salt,api_key FROM users WHERE username = ?";
	rows, err := repo.DB.QueryContext(ctx, stmt, username);
	helper.PanicIfError(err);
	defer rows.Close();

	var user domain.User;
	if rows.Next() {
		err := rows.Scan(&user.Id,&user.Username,&user.Password,&user.Salt,&user.ApiKey);
		helper.PanicIfError(err);
		return user,true
	} else {
		return user,false
	}
}



func (repo *AuthRepositoryImpl) FindBySession(ctx context.Context,sessionKey string) (web.UserResponse,error) {

	var user web.UserResponse;
	
	stmt := "SELECT id, username, session_key, session_expires_at, api_key FROM users WHERE session_key = ? limit 1";
	rows, err := repo.DB.QueryContext(ctx, stmt, sessionKey);

	
	if err != nil {
		fmt.Println("error repo query",err.Error())
		return user,err;
	}
	defer rows.Close();


	if rows.Next(){
		err = rows.Scan(&user.Id,&user.Username,&user.SessionKey,&user.SessionExpiresAt,&user.ApiKey);
		if err != nil {
            fmt.Println("Error scanning row:", err.Error())
            return user, err
        }
		return user,nil;
	} 
	return user, errors.New("user not found")

	
}

func (repo *AuthRepositoryImpl) DestroySession(ctx context.Context,sessionKey string) (int64,error){
	stmt := "UPDATE users SET session_key = ? WHERE session_key = ?";
	result,err := repo.DB.ExecContext(ctx,stmt,nil,sessionKey);
	if err != nil {
		return 0,err
	}

	row, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected:", err.Error())
		return 0,errors.New("logout server error");
	}

	return row,nil
}



