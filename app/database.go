package app

import (
	"database/sql"
	"shinkansen_rest_api/config"
	"shinkansen_rest_api/helper"
	"time"
)

func NewDB() *sql.DB{
	dsn := config.GetDsnEnv();

	db, err := sql.Open("mysql",dsn);
	helper.PanicIfError(err);

	db.SetMaxIdleConns(10);
	db.SetMaxIdleConns(50);
	db.SetConnMaxLifetime(time.Minute * 60);
	db.SetConnMaxIdleTime(time.Minute * 10);

	return db;
}