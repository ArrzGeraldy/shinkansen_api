package service

import (
	"context"
	"shinkansen_rest_api/exception"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/models/domain"
	"shinkansen_rest_api/repository"
	"strconv"
)

type StationService interface {
	GetStations(ctx context.Context, pageNumberStr string,pageSize int) ([]domain.Station,int,int,int,error)
	GetStationById(ctx context.Context,stationIdStr string) (domain.Station);

}

type StationServiceImpl struct{
	Repo repository.StationRepository
}

func NewStationService(repo repository.StationRepository) StationService{
	return &StationServiceImpl{
		Repo: repo,
	}
}

func (service *StationServiceImpl) GetStations(ctx context.Context,pageNumberStr string,pageSize int) ([]domain.Station,int,int,int,error){

	pageNumber, err := helper.PageQueryParamsValidation(pageNumberStr);

	if err != nil {
		panic(exception.ErrParameter);
	}

	stations := service.Repo.GetStations(ctx,pageNumber,pageSize);

	totalItems,err := service.Repo.CountStations()
	helper.PanicIfError(err);

	mod := totalItems % pageSize;
	totalPage := totalItems / pageSize;

	if mod > 0 {
		totalPage++
	}

	return stations,totalItems,totalPage,pageNumber,nil;
}


func (service *StationServiceImpl) GetStationById(ctx context.Context,stationIdStr string) (domain.Station){

	var station domain.Station;

	stationId,err := strconv.Atoi(stationIdStr);
	if err != nil {
		panic(exception.ErrParameter)
	}

	station,err =  service.Repo.GetStationById(ctx,stationId);
	helper.PanicIfError(err);

	return station;

}

