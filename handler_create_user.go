package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name     string `json:"name"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	newUser, err := apiCfg.dbClient.CreateUser(params.Name, params.Password, params.Email, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(r.Method)
	respondWithJSON(w, http.StatusCreated, newUser)
}
