package main

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
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

// GetUser
// @Summary get user
// @Description Get user with matching passport
// @Tags users
// @Accept json
// @Produce json
// @Param query body models.UserPassport true "passport data"
// @Success 200 {object} models.User
// @Failure 400 {object} JSONResponse
// @Router /users [post]
func (app *TrackerApp) GetUser(w http.ResponseWriter, r *http.Request) {
	var userPassport models.UserPassport
	err := readJSONPayload(w, r, &userPassport)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	err = validatePassport(userPassport.PassportNumber)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.DB.SelectUserByPassport(userPassport.PassportNumber)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	writeJSONResponse(w, http.StatusOK, user)
}

// GetUserById
// @Summary get user by id
// @Description Get user by their internal id
// @Tags users
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} models.User
// @Failure 400 {object} JSONResponse
// @Router /users/{id} [get]
func (app *TrackerApp) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		writeJSONError(w, errors.New("bad user id"), http.StatusBadRequest)
		return
	}

	user, err := app.DB.SelectUserById(userId)
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

	err = validatePassport(user.Passport)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	// check whether a user with this passport already exists
	foundUser, err := app.DB.SelectUserByPassport(user.Passport)
	if !errors.Is(err, sql.ErrNoRows) {
		resp := JSONResponse{
			Error:   true,
			Message: "user with passport " + user.Passport + " already exists",
			Data:    foundUser,
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
		Message: "user " + user.Passport + " has been created",
	}
	writeJSONResponse(w, http.StatusCreated, resp)
}

// UpdateUser
// @Summary update user
// @Description Update data on existing user provided their passport
// @Tags users
// @Accept json
// @Produce json
// @Param query body models.User true "user data"
// @Success 200 {object} JSONResponse
// @Failure 400 {object} JSONResponse
// @Router /users [patch]
func (app *TrackerApp) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := readJSONPayload(w, r, &user)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	err = app.DB.UpdateUser(user)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "user " + user.Passport + " has been updated",
	}
	writeJSONResponse(w, http.StatusOK, resp)

}

// DeleteUser
// @Summary delete user
// @Description Delete user provided their passport
// @Tags users
// @Accept json
// @Produce json
// @Param query body models.UserPassport true "user data"
// @Success 200 {object} JSONResponse
// @Failure 400 {object} JSONResponse
// @Router /users [delete]
func (app *TrackerApp) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userPassport models.UserPassport
	err := readJSONPayload(w, r, &userPassport)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	err = validatePassport(userPassport.PassportNumber)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	err = app.DB.DeleteUserByPassport(userPassport.PassportNumber)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "user " + userPassport.PassportNumber + " has been deleted",
	}
	writeJSONResponse(w, http.StatusOK, resp)
}

// GetAllTasks
// @Summary get all tasks
// @Description Return a list of all tasks
// @Tags tasks
// @Produce json
// @Success 200 {object} models.ResponseTasks
// @Failure 500 {object} JSONResponse
// @Router /tasks [get]
func (app *TrackerApp) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.DB.SelectAllTasks()
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	tasksResponse := models.ResponseTasks{Tasks: tasks}
	writeJSONResponse(w, http.StatusOK, tasksResponse)
}

// StartTask
// @Summary start task
// @Description Start a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param query body models.RequestNewTask true "task data"
// @Success 201 {object} JSONResponse
// @Failure 400 {object} JSONResponse
// @Failure 500 {object} JSONResponse
// @Router /tasks [post]
func (app *TrackerApp) StartTask(w http.ResponseWriter, r *http.Request) {
	var task models.RequestNewTask

	err := readJSONPayload(w, r, &task)
	if err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	taskId, err := app.DB.StartTask(task)
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "new task with id " + strconv.Itoa(taskId) + " has been started",
	}
	writeJSONResponse(w, http.StatusCreated, resp)
}

// FinishTask
// @Summary finish task
// @Description Finish the task
// @Tags tasks
// @Produce json
// @Param id path int true "task id"
// @Success 200 {object} JSONResponse
// @Failure 400 {object} JSONResponse
// @Router /tasks/{id} [post]
func (app *TrackerApp) FinishTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		writeJSONError(w, errors.New("bad task id"), http.StatusBadRequest)
		return
	}

	err = app.DB.FinishTask(taskId)
	if err != nil {
		writeJSONError(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "task " + strconv.Itoa(taskId) + " has been finished",
	}
	writeJSONResponse(w, http.StatusOK, resp)
}

// DeleteTask
// @Summary delete task
// @Description Delete the task
// @Tags tasks
// @Produce json
// @Param id path int true "task id"
// @Success 200 {object} JSONResponse
// @Failure 400 {object} JSONResponse
// @Router /tasks/{id} [delete]
func (app *TrackerApp) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		writeJSONError(w, errors.New("bad task id"), http.StatusBadRequest)
		return
	}

	err = app.DB.DeleteTask(taskId)
	if err != nil {
		writeJSONError(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "task " + strconv.Itoa(taskId) + " has been deleted",
	}
	writeJSONResponse(w, http.StatusOK, resp)
}
