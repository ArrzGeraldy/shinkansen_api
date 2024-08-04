package repository

import (
	"context"
	"database/sql"
	"shinkansen_rest_api/exception"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/models/domain"
)

type StationRepository interface {
	GetStations(ctx context.Context,pageNumber int,pageSize int) []domain.Station;
	GetStationById(ctx context.Context,id int) (domain.Station,error);
	CountStations() (int,error);
}

type StationRepositoryImpl struct{
	DB *sql.DB
}

func NewStationRepository(db *sql.DB) StationRepository{
	return &StationRepositoryImpl{
		DB: db,
	}
}

func (repo *StationRepositoryImpl) GetStations(ctx context.Context,pageNumber int,pageSize int) []domain.Station{
	stmt := "SELECT id,station_name,shinkansen_line,year,prefecture,distance_from_tokyo,company FROM shinkansen_stations LIMIT ? OFFSET ?";

	offset := (pageNumber-1) * pageSize;

	rows,err := repo.DB.QueryContext(ctx,stmt,pageSize,offset);
	helper.PanicIfError(err);

	defer rows.Close();

	var stations []domain.Station;

	for rows.Next(){
		var station domain.Station;

		err = rows.Scan(&station.Id,&station.StationName,&station.ShinkansenLine,&station.Year,&station.Prefecture,&station.DistanceFromTokyo,&station.Company);
		helper.PanicIfError(err);

		stations = append(stations, station);
	}

	return stations;
}

func (repo *StationRepositoryImpl) GetStationById(ctx context.Context,id int) (domain.Station,error){
	stmt := "SELECT id,station_name,shinkansen_line,year,prefecture,distance_from_tokyo,company FROM shinkansen_stations WHERE id = ?";

	rows,err := repo.DB.QueryContext(ctx,stmt,id);
	helper.PanicIfError(err);
	defer rows.Close();

	var station domain.Station;
	if rows.Next(){
		err = rows.Scan(&station.Id,&station.StationName,&station.ShinkansenLine,&station.Year,&station.Prefecture,&station.DistanceFromTokyo,&station.Company);
		helper.PanicIfError(err);

		return station,nil;
	} else {
		return station,exception.ErrNotFound;
	}

}

func (repo *StationRepositoryImpl) CountStations() (int,error){

	rows := repo.DB.QueryRow("SELECT COUNT(*) FROM shinkansen_stations");

	var total int;
	err := rows.Scan(&total);

	if err != nil {
		helper.PanicIfError(err);
	}

	return	total,nil;

}