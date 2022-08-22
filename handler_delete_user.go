package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting user")
	userEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	if userEmail == "" {
		fmt.Println("error 1, email empty")
		respondWithError(w, http.StatusBadRequest, errors.New("no userEmail has been provided to handlerUpdateUser"))
		return
	}
	err := apiCfg.dbClient.DeleteUser(userEmail)
	if err != nil {
		fmt.Println("Error 2, bad parameters")
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("If json body is empty, then you're good!")
	respondWithJSON(w, http.StatusOK, struct{}{})
}
