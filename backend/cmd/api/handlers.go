package main

import (
	"backend/internal/models"
	"net/http"
)

// AllUsers
// @Summary All users
// @Description Return a list of all users
// @Tags users
// @Produce json
// @Success 200 {object} models.ResponseUsers
// @Failure 500 {object} JSONResponse
// @Router /users [get]
func (app *TrackerApp) AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.DB.AllUsers()
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	usersResponse := models.ResponseUsers{Users: users}
	writeJSONResponse(w, http.StatusOK, usersResponse)
}
