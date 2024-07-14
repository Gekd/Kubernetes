package handlers

import (
	"encoding/json"
	"identity-server/utilities"
	"identity-server/variables"
	"net/http"
	"os"
	"time"

	"github.com/go-oauth2/oauth2/v4/server"
)

func IntrospectionHandler(responseWriter http.ResponseWriter, request *http.Request, authorizationServer *server.Server) {
	if variables.DumpRequestsAndResponses {
		_ = utilities.LogRequests(os.Stdout, "test", request) // Ignore the error
	}
	token, err := authorizationServer.ValidationBearerToken(request)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	}
	e := json.NewEncoder(responseWriter)
	e.SetIndent("", "  ")
	e.Encode(data)
}
