package api

import (
	"database/sql"
	"net/http"
	"shinkansen_rest_api/controller"
	"shinkansen_rest_api/middleware"
	"shinkansen_rest_api/repository"
	"shinkansen_rest_api/service"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB) {
	stationRepository := repository.NewStationRepository(db)
	stationService := service.NewStationService(stationRepository)
	stationController := controller.NewStationController(stationService)

	mux.Handle("/api/shinkansen/stations/", middleware.ApiAuth(http.HandlerFunc(stationController.GetStations)))
	mux.Handle("/api/shinkansen/stations/{id}", middleware.ApiAuth(http.HandlerFunc(stationController.GetStationById)))

}