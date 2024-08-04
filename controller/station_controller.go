package controller

import (
	"net/http"
	"shinkansen_rest_api/helper"
	"shinkansen_rest_api/models/web"
	"shinkansen_rest_api/service"
)

type StationController interface {
	GetStations(writer http.ResponseWriter,req *http.Request)
	GetStationById(writer http.ResponseWriter,req *http.Request)
}

type StationControllerImpl struct{
	Service service.StationService
}

func NewStationController(service service.StationService) StationController{
	return &StationControllerImpl{
		Service: service,
	}
}


func(controller *StationControllerImpl) GetStations(writer http.ResponseWriter,req *http.Request){
	pageNumberStr := req.URL.Query().Get("p");
	pageSize := 10;
	
	stations,totalItems,totalPage,pageNumber,err := controller.Service.GetStations(req.Context(),pageNumberStr,pageSize);
	helper.PanicIfError(err);

	response := web.WebResponsePagination{
		Code: 200,
		Status: "OK",
		Data: stations,
		Pagination: web.Pagination{
			PageSize: pageSize,
			CurrentPage: pageNumber,
			TotalPage: totalPage,
			TotalItems: totalItems,
		},
	}

	helper.ToResponseJson(writer,response);
}

func (controller *StationControllerImpl) GetStationById(writer http.ResponseWriter,req *http.Request){
	stationIdStr := helper.GetParamURL(req,"/api/shinkansen/stations/");

	station := controller.Service.GetStationById(req.Context(),stationIdStr);

	response := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: station,
	}

	helper.ToResponseJson(writer,response);
}