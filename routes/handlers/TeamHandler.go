package handlers

import (
	"ProgettoDB/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TeamsHandler struct {
	Ctrl controllers.TeamController
}

func (t TeamsHandler) GETAllTeams(writer http.ResponseWriter, _ *http.Request) {
	if champs, err := t.Ctrl.GetAllTeams(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, champs)
	}
}

func (t TeamsHandler) GETTeamsWithoutParticipationByYear(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	strYear := params["year"]

	if strYear == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param name"))
		return
	}
	if year, err := strconv.Atoi(strYear); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("wrong param year, it has to be a number"))
		return
	} else {
		if champs, err2 := t.Ctrl.GetTeamsWithoutParticipationByYear(uint(year)); err2 != nil {
			respondError(writer, http.StatusInternalServerError, err2)
		} else {
			respondJSON(writer, http.StatusOK, champs)
		}
	}

}
