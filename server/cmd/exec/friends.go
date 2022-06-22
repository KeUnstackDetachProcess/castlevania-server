package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getFriendsList(rw http.ResponseWriter, r *http.Request) {

}

func sendFriendRequest(rw http.ResponseWriter, r *http.Request) {

}

func removeUserFromFriends(rw http.ResponseWriter, r *http.Request) {

}

func InitializeFriendsRoutes(router *mux.Router) {
	// retrieve all user friends
	router.HandleFunc("/api/user/#:client/friends", getFriendsList).Methods("GET")
	// retrieve all user friends
	router.HandleFunc("/api/user/#:client/friends/add/#:target", sendFriendRequest).Methods("POST")
	// remove an user from friends
	router.HandleFunc("/api/user/#:client/friends/remove/#:target", removeUserFromFriends).Methods("POST")
}
