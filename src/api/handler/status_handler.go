package handler

import (
	"encoding/json"
	"fmt"
	"github.com/josewenzel/flagify/src/api/security"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func HandleGetStatus(responseWriter http.ResponseWriter, request *http.Request) {
	if authorizeRequest := security.AuthorizeRequest(request); !authorizeRequest.Valid {
		StatusRespondWith(http.StatusUnauthorized, responseWriter)
		return
	}

	StatusRespondWith(http.StatusOK, responseWriter)
}

func StatusRespondWith(statusCode int, responseWriter http.ResponseWriter) {
	responseWriter.WriteHeader(statusCode)
	response := &StatusResponse{Status: http.StatusText(statusCode)}
	jsonResponse, _ := json.Marshal(response)
	fmt.Fprintf(responseWriter, string(jsonResponse))
}
