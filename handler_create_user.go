package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Parameters needed to create user in db.json file
func (apiCfg apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user")
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

	newUser, err := apiCfg.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusCreated, newUser)
}
