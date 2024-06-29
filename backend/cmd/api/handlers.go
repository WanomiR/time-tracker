package main

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"net/http"
)

// GetAllUsers
// @Summary get all users
// @Description Return a list of all users
// @Tags users
// @Produce json
// @Success 200 {object} models.ResponseUsers
// @Failure 500 {object} JSONResponse
// @Router /users [get]
func (app *TrackerApp) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.DB.SelectAllUsers()
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	usersResponse := models.ResponseUsers{Users: users}
	writeJSONResponse(w, http.StatusOK, usersResponse)
}

// GetUserByPassport
// @Summary user by passport
// @Description Returns user with matching passport
// @Tags users
// @Accept json
// @Produce json
// @Param query body models.UserPassport true "passport data"
// @Success 200 {object} models.User
// @Failure 400 {object} JSONResponse
// @Router /users [post]
func (app *TrackerApp) GetUserByPassport(w http.ResponseWriter, r *http.Request) {
	var userPassport models.UserPassport
	err := readJSONPayload(w, r, &userPassport)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.DB.SelectUserByPassport(userPassport.Series, userPassport.Number)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	writeJSONResponse(w, http.StatusOK, user)
}

// AddUser
// @Summary add new user
// @Description Add new user
// @Tags users
// @Accept json
// @Produce json
// @Param query body models.User true "user data"
// @Success 201 {object} JSONResponse
// @Failure 400 {object} JSONResponse
// @Failure 409 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Router /users [put]
func (app *TrackerApp) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := readJSONPayload(w, r, &user)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	// check whether a user with this passport already exists
	receivedUser, err := app.DB.SelectUserByPassport(user.PassportSeries, user.PassportNumber)
	if !errors.Is(err, sql.ErrNoRows) {
		resp := JSONResponse{
			Error:   true,
			Message: "User with this passport already exists",
			Data:    receivedUser,
		}
		writeJSONResponse(w, http.StatusConflict, resp)
		return
	}

	err = app.DB.InsertUser(user)
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "user created",
	}
	writeJSONResponse(w, http.StatusCreated, resp)
}
