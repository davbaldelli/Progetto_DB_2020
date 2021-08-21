package handlers

import (
	"ProgettoDB/controllers"
	"net/http"
)

type DriversHandler struct {
	Ctrl controllers.DriversController
}

func (d DriversHandler) GETAllDrivers(writer http.ResponseWriter, request *http.Request) {
	if drivers, err := d.Ctrl.GetAllDrivers(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, drivers)
	}
}
