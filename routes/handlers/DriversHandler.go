package handlers

import (
	"ProgettoDB/controllers"
	"ProgettoDB/models"
	"encoding/json"
	"fmt"
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

func (d DriversHandler) POSTNewDriver(writer http.ResponseWriter, request *http.Request) {
	driver := models.Driver{}

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&driver); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := d.Ctrl.InsertDriver(driver); err != nil {
		respondError(writer, http.StatusInternalServerError, fmt.Errorf("cannot insert new entity: %v ", err))
		return
	}

	respondJSON(writer, http.StatusCreated, driver)
}
