package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Get user post
func (apiCfg apiConfig) handlerGetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("retrieving post")
	userEmail := r.URL.Query().Get("userEmail")
	if userEmail == "" {
		fmt.Println("error, get request made for non-existent userID")
		respondWithError(w, http.StatusBadRequest, errors.New("cannot make request without userEmail"))
		return
	}
	posts, err := apiCfg.dbClient.GetPosts(userEmail)
	if err != nil {
		fmt.Println("Bad parameters")
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("Status OK")
	respondWithJSON(w, http.StatusOK, posts)
}
