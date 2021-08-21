package handlers

import (
	"ProgettoDB/controllers"
	"net/http"
)

type ClassesHandler struct {
	Ctrl controllers.ClassesController
}

func (c ClassesHandler) GETAllClasses(writer http.ResponseWriter, request *http.Request) {
	if classes, err := c.Ctrl.GetAllClasses(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, classes)
	}
}
