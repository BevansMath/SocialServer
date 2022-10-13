package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Delete user
func (apiCfg apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	userEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	fmt.Println(userEmail)
	if userEmail == "" {
		fmt.Println("error 1, email empty")
		respondWithError(w, http.StatusBadRequest, errors.New("no userEmail has been provided to handlerUpdateUser"))
		return
	}
	err := apiCfg.dbClient.DeleteUser(userEmail)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(apiCfg.dbClient.DeleteUser(userEmail))
	respondWithJSON(w, http.StatusOK, struct{}{})
}
