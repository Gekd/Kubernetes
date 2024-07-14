package handlers

import (
	"identity-server/utilities"
	"identity-server/variables"
	"net/http"
	"net/url"
	"os"

	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-session/session"
)

func AuthorizeHandler(w http.ResponseWriter, r *http.Request, authorizationServer *server.Server) {
	if variables.DumpRequestsAndResponses {
		utilities.LogRequests(os.Stdout, "authorize", r)
	}

	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	r.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = authorizationServer.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
