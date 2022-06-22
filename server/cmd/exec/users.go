package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func usersRequestList(rw http.ResponseWriter, r *http.Request) {

}

func userRequestCreation(rw http.ResponseWriter, r *http.Request) {

}

func userRequestSelection(rw http.ResponseWriter, r *http.Request) {

}

func userRequestConnection(rw http.ResponseWriter, r *http.Request) {

}

func userRequestDeletion(rw http.ResponseWriter, r *http.Request) {

}

func InitializeUsersRoutes(router *mux.Router) {
	// retrieve all users
	router.HandleFunc("/api/users", usersRequestList).Methods("GET")
	// create new users
	router.HandleFunc("/api/users/create", userRequestCreation).Methods("POST")
	// retrieve user data
	router.HandleFunc("/api/users/#:tag", userRequestSelection).Methods("GET")
	// connect to a new or existing peer-to-peer session with the requested user
	router.HandleFunc("/api/users/#:client/connect", userRequestConnection).Methods("POST")
	// delete an user's account, only the user can delete their own account
	router.HandleFunc("/api/users/#:tag/delete", userRequestDeletion).Methods("POST")
}
