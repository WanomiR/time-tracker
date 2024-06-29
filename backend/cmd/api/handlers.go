package main

import "net/http"

func (app *TrackerApp) AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.DB.AllUsers()
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, http.StatusOK, users)
}
