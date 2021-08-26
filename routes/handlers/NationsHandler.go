package handlers

import (
	"ProgettoDB/controllers"
	"net/http"
)

type NationsHandler struct {
	Ctrl controllers.NationsController
}

func (n NationsHandler) GETAllNations(writer http.ResponseWriter, _ *http.Request){
	if nations, err := n.Ctrl.GetAllNations(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, nations)
	}
}