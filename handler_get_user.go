package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Get user information
func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating User")
	userEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	if userEmail == "" {
		fmt.Println("No email specified")
		respondWithError(w, http.StatusBadRequest, errors.New("no user email provided to handlerGetUser"))
		return
	}
	user, err := apiCfg.dbClient.GetUser(userEmail)
	if err != nil {
		fmt.Println("Bad Parameters")
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("All good!")
	respondWithJSON(w, http.StatusOK, user)
}
