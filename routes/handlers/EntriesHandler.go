package handlers

import (
	"ProgettoDB/controllers"
	"ProgettoDB/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EntriesHandler struct {
	Ctrl controllers.EntriesController
}

func (e EntriesHandler) GETChampionshipEntries(writer http.ResponseWriter, request *http.Request) {
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
		if entries, err2 := e.Ctrl.GetChampionshipEntryList(models.Championship{Name: champ, Year: uint(year)}); err2 != nil {
			if err2.Error() == "not found" {
				respondError(writer, http.StatusNotFound, err)
			} else {
				respondError(writer, http.StatusInternalServerError, err)
			}
		} else {
			respondJSON(writer, http.StatusOK, entries)
		}
	}
}

func (e EntriesHandler) GETEntryByRaceNumber(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	champ := params["name"]
	strYear := params["year"]
	strNumber := params["number"]

	if champ == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param name"))
		return
	}

	if strYear == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param year"))
		return
	}

	if strYear == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param number"))
		return
	}

	if year, err := strconv.Atoi(strYear); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("wrong param year, it has to be a number"))
		return
	} else {
		if number, err3 := strconv.Atoi(strNumber); err3 != nil {
			respondError(writer, http.StatusBadRequest, fmt.Errorf("wrong param number, it has to be a number"))
			return
		} else {
			if entries, err2 := e.Ctrl.GetEntryByRaceNumber(models.Championship{Name: champ, Year: uint(year)}, uint(number)); err2 != nil {
				if err2.Error() == "not found" {
					respondError(writer, http.StatusNotFound, err)
				} else {
					respondError(writer, http.StatusInternalServerError, err)
				}
			} else {
				respondJSON(writer, http.StatusOK, entries)
			}
		}

	}
}
