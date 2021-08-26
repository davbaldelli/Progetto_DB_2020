package handlers

import (
	"ProgettoDB/controllers"
	"ProgettoDB/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CarHandler struct {
	Ctrl controllers.CarController
}

func (c CarHandler) GETChampionshipCars(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	champ := params["name"]
	strYear := params["year"]

	if champ == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param name"))
		return
	}

	if strYear == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param year"))
		return
	}

	if year, err := strconv.Atoi(strYear); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("wrong param year, it has to be a number"))
		return
	} else {
		if cars, err2 := c.Ctrl.GetChampionshipCars(models.Championship{Name: champ, Year: uint(year)}); err2 != nil {
			if err2.Error() == "not found" {
				respondError(writer, http.StatusNotFound, err2)
			} else {
				respondError(writer, http.StatusInternalServerError, err2)
			}
		} else {
			respondJSON(writer, http.StatusOK, cars)
		}
	}

}

func (c CarHandler) GETDriverCarsOnTrack(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	driver := params["driver"]
	track := params["track"]

	if driver == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param driver"))
		return
	}

	if track == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param track"))
		return
	}

	if cars, err := c.Ctrl.GetDriverCarOnCircuit(models.Driver{CF: driver}, models.Track{Name: track}); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, cars)
	}
}
