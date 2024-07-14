package handlers

import (
	"identity-server/utilities"
	"identity-server/variables"
	"net/http"
	"os"

	"github.com/go-session/session"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if variables.DumpRequestsAndResponses {
		_ = utilities.LogRequests(os.Stdout, "auth", r) // Ignore the error
	}
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	utilities.OutputHTML(w, r, "static/auth.html")
}
