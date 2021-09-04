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
	class := params["class"]

	if class == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param class"))
		return
	}

	if races, err := r.Ctrl.GetRacesByClass(class); err != nil {
		if err.Error() == "not found" {
			respondError(writer, http.StatusNotFound, err)
		} else {
			respondError(writer, http.StatusInternalServerError, err)
		}
	} else {
		respondJSON(writer, http.StatusOK, races)
	}
}

func (r RacesHandler) GETAllRaces(writer http.ResponseWriter, _ *http.Request) {
	if races, err := r.Ctrl.GetAllRaces(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, races)
	}
}

func (r RacesHandler) GETRacesByTeam(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	team := params["team"]

	if team == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param team"))
		return
	}

	if races, err := r.Ctrl.GetRacesByTeam(team); err != nil {
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
	nation := params["nation"]

	if nation == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param nation"))
		return
	}

	if races, err := r.Ctrl.GetDriversRacesByNationality(nation); err != nil {
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
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param champ"))
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
				respondError(writer, http.StatusNotFound, err2)
			} else {
				respondError(writer, http.StatusInternalServerError, err2)
			}
		} else {
			respondJSON(writer, http.StatusOK, races)
		}
	}

}

func (r RacesHandler) GETRacesResult(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	champ := params["champ"]
	strYear := params["year"]
	race := params["race"]

	if race == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param race"))
		return
	}

	if champ == "" {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("missing param champ"))
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
		if races, err2 := r.Ctrl.GetRaceResult(models.Race{Name: race, Championship: models.Championship{Name: champ, Year: uint(year)}}); err2 != nil {
			if err2.Error() == "not found" {
				respondError(writer, http.StatusNotFound, err2)
			} else {
				respondError(writer, http.StatusInternalServerError, err2)
			}
		} else {
			respondJSON(writer, http.StatusOK, races)
		}
	}

}
