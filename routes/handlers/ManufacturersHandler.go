package handlers

import (
	"ProgettoDB/controllers"
	"net/http"
)

type ManufacturersHandler struct {
	Ctrl controllers.ManufacturersController
}

func (m ManufacturersHandler) GETAllManufacturers(writer http.ResponseWriter, _ *http.Request) {
	if manufacturers, err := m.Ctrl.GetAllManufacturers(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, manufacturers)
	}
}
