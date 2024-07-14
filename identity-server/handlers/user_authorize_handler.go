package handlers

import (
	"identity-server/utilities"
	"identity-server/variables"
	"net/http"
	"os"

	"github.com/go-session/session"
)

func UserAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	if variables.DumpRequestsAndResponses {
		_ = utilities.LogRequests(os.Stdout, "userAuthorizeHandler", r) // Ignore the error
	}
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		store.Set("ReturnUri", r.Form)
		store.Save()

		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	userID = uid.(string)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
