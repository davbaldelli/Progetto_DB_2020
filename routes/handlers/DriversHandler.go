package handlers

import (
	"ProgettoDB/controllers"
	"net/http"
)

type DriversHandler struct {
	Ctrl controllers.DriversController
}

func (d DriversHandler) GETAllDrivers(writer http.ResponseWriter, _ *http.Request) {
	if drivers, err := d.Ctrl.GetAllDrivers(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, drivers)
	}
}

func (d DriversHandler) GETFiveDriversWithMoreRaces(writer http.ResponseWriter, _ *http.Request) {
	if drivers, err := d.Ctrl.GetFiveDriversWithMoreRaces(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, drivers)
	}
}
