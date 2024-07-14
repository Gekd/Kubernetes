package handlers

import (
	"identity-server/src/utilities"
	"identity-server/src/variables"
	"net/http"
	"os"

	"github.com/go-oauth2/oauth2/v4/server"
)

func TokenHandler(responseWriter http.ResponseWriter, request *http.Request, authorizationServer *server.Server) {
	if variables.DumpRequestsAndResponses {
		_ = utilities.LogRequests(os.Stdout, "token", request) // Ignore the error
	}

	err := authorizationServer.HandleTokenRequest(responseWriter, request)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
