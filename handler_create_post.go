package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define post required fields
func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	newPost, err := apiCfg.dbClient.CreatePost(params.UserEmail, params.Text)
	if err != nil {
		fmt.Println("bad parameters, post not created")
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("parameters OK, creating post")
	respondWithJSON(w, http.StatusCreated, newPost)
}
