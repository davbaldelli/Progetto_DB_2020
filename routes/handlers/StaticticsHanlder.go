package handlers

import (
	"ProgettoDB/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type StatisticsHandler struct {
	Ctrl controllers.StatisticsController
}

func (s StatisticsHandler) GETBrandsCarsUsage(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	brand := params["brand"]

	if brand == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param brand"))
		return
	}

	if statistics, err := s.Ctrl.GetBrandCarsUsage(brand); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, statistics)
	}
}

func (s StatisticsHandler) GETTrackLayoutsUsage(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	track := params["track"]

	if track == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param track"))
		return
	}

	if statistics, err := s.Ctrl.GetTrackLayoutsUsage(track); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, statistics)
	}
}
