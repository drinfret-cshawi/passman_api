package controllers

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"passman_api/db"
	"strconv"
)

type JsonUserResponse struct {
	Type string    `json:"type"`
	Data []db.User `json:"data"`
}

func GetUsers(w http.ResponseWriter, _ *http.Request) {

	users, err := db.GetUsers()
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}
	var response = JsonUserResponse{Type: "success", Data: users}
	handleJSONResponse(response, w)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("Controller: GetUserById", id)

	user, err := db.GetUserById(id)
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}
	var users []db.User
	users = append(users, user)
	var response = JsonUserResponse{Type: "success", Data: users}
	handleJSONResponse(response, w)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	fullname := r.FormValue("fullname")
	password := r.FormValue("password")
	email := r.FormValue("email")

	fmt.Println("Controller: AddUser", username, fullname, password, email)

	if username == "" || password == "" {
		handleError(errors.New("username and password cannot be empty"), w)
		return
	}
	id, err := db.AddUser(username, fullname, password, email)
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}

	var response = JsonResponse{Type: "success", Message: strconv.Itoa(id)}
	handleJSONResponse(response, w)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println("Controller: DeleteUser", id)

	nRows, err := db.DeleteUser(id)
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}

	var response = JsonResponse{Type: "success", Message: strconv.Itoa(nRows)}
	handleJSONResponse(response, w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	nRows, _ := strconv.Atoi(params["id"])

	username := r.FormValue("username")
	fullname := r.FormValue("fullname")
	password := r.FormValue("password")
	email := r.FormValue("email")

	fmt.Println("Controller: UpdateUser", nRows, username, fullname, password, email)

	if username == "" || password == "" {
		handleError(errors.New("username and password cannot be empty"), w)
		return
	}
	nRows, err := db.UpdateUser(nRows, username, fullname, password, email)
	if err != nil {
		//TODO: don't send back db error directly
		handleError(err, w)
		return
	}
	if nRows == 0 {
		handleError(errors.New("user not found"), w)
		return
	}

	var response = JsonResponse{Type: "success", Message: strconv.Itoa(nRows)}
	handleJSONResponse(response, w)
}
