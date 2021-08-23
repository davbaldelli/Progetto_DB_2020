package handlers

import (
	"ProgettoDB/controllers"
	"net/http"
)

type TracksHandler struct {
	Ctrl controllers.TracksController
}

func (t TracksHandler) GETAllTracks(writer http.ResponseWriter, _ *http.Request) {
	if tracks, err := t.Ctrl.GetAllTracks(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, tracks)
	}
}
