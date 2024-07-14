package handlers

import (
	"identity-server/src/utilities"
	"identity-server/src/variables"
	"net/http"
	"os"

	"github.com/go-session/session"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if variables.DumpRequestsAndResponses {
		_ = utilities.LogRequests(os.Stdout, "login", r) // Ignore the error
	}
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		if r.Form == nil {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		store.Set("LoggedInUserID", r.Form.Get("username"))
		store.Save()

		w.Header().Set("Location", "/auth")
		w.WriteHeader(http.StatusFound)
		return
	}
	utilities.OutputHTML(w, r, "src/static/login.html")
}
