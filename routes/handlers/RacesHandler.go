package handlers

import (
	"ProgettoDB/controllers"
	"ProgettoDB/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type RacesHandler struct {
	Ctrl controllers.RacesController
}

func (r RacesHandler) GETRacesByClass(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params["class"]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param class"))
		return
	}

	if races, err := r.Ctrl.GetRacesByClass(param); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, races)
	}
}

func (r RacesHandler) GETRacesByTeam(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params["team"]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param team"))
		return
	}

	if races, err := r.Ctrl.GetRacesByTeam(param); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, races)
	}
}

func (r RacesHandler) GETRacesByDriversNation(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params["nation"]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param nation"))
		return
	}

	if races, err := r.Ctrl.GetDriversRacesByNationality(param); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, races)
	}
}

func (r RacesHandler) GETRacesByChampionship(writer http.ResponseWriter, request *http.Request) {
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
		if races, err2 := r.Ctrl.GetChampionshipRaces(models.Championship{Name: champ, Year: uint(year)}); err2 != nil {
			if err2.Error() == "not found" {
				respondError(writer, http.StatusNotFound, err)
			} else {
				respondError(writer, http.StatusInternalServerError, err)
			}
		} else {
			respondJSON(writer, http.StatusOK, races)
		}
	}

}
