package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"passman_api/db"
	"strconv"
)

type JsonPasswordResponse struct {
	Type string        `json:"type"`
	Data []db.Password `json:"data"`
}

func GetPasswordsForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["user_id"])
	fmt.Println("Controller: GetPasswordsForUser")

	passwords, err := db.GetPasswordsForUser(userId)
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}
	var response = JsonPasswordResponse{Type: "success", Data: passwords}
	handleJSONResponse(response, w)

}

func AddPassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["user_id"])
	site := r.FormValue("site")
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("Controller: AddPassword", userId, site, username, password)

	if site == "" || password == "" {
		handleError(errors.New("site and password cannot be empty"), w)
		return
	}
	id, err := db.AddPassword(userId, site, username, password)
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}

	var response = JsonResponse{Type: "success", Message: strconv.Itoa(id)}
	handleJSONResponse(response, w)
}

func DeletePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["user_id"])
	site := r.FormValue("site")
	username := r.FormValue("username")

	fmt.Println("Controller: DeletePassword", userId, site, username)

	if site == "" {
		handleError(errors.New("site cannot be empty"), w)
		return
	}
	nRows, err := db.DeletePassword(userId, site, username)
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}

	var response = JsonResponse{Type: "success", Message: strconv.Itoa(nRows)}
	handleJSONResponse(response, w)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["user_id"])
	site := r.FormValue("site")
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println("Controller: UpdatePassword", userId, site, username, password)

	if site == "" {
		handleError(errors.New("site cannot be empty"), w)
		return
	}
	nRows, err := db.UpdatePassword(userId, site, username, password)
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}

	var response = JsonResponse{Type: "success", Message: strconv.Itoa(nRows)}
	handleJSONResponse(response, w)
}
