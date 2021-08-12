package handlers

import (
	"ProgettoDB/controllers"
	"ProgettoDB/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ChampionshipsHandler struct {
	Ctrl controllers.ChampionshipsController
}

func (c ChampionshipsHandler) GETAllChampionships(writer http.ResponseWriter, request *http.Request) {
	if champs, err := c.Ctrl.GetAllChampionships(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, champs)
	}
}

func (c ChampionshipsHandler) GETDriverChampionships(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params["cf"]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param cf"))
		return
	}

	if champs, err := c.Ctrl.GetDriverChampionships(models.Driver{CF: param}); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, champs)
	}

}

func (c ChampionshipsHandler) GETTeamChampionships(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params["team"]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param team"))
		return
	}

	if champs, err := c.Ctrl.GetChampionshipsByTeam(models.Team{Name: param}); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, champs)
	}

}

func (c ChampionshipsHandler) GETDriversChampionshipsByNation(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	param := params["nation"]

	if param == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param nation"))
		return
	}

	if champs, err := c.Ctrl.GetDriversChampionshipsByNationality(param); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, champs)
	}

}
