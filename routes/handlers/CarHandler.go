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
		if statistics, err2 := c.Ctrl.GetChampionshipCars(models.Championship{Name: champ, Year: uint(year)}); err2 != nil {
			if err2.Error() == "not found" {
				respondError(writer, http.StatusNotFound, err2)
			} else {
				respondError(writer, http.StatusInternalServerError, err2)
			}
		} else {
			respondJSON(writer, http.StatusOK, statistics)
		}
	}

}
