package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	userEmail := strings.TrimPrefix(r.URL.Path, "/users/")
	type parameters struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	if userEmail == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("no email provided to handleUpdateUser"))
		return
	}
	newUser, err := apiCfg.dbClient.UpdateUser(userEmail, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, newUser)
}
