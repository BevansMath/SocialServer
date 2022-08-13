package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) { // Create user endpoint
	fmt.Println(r.Method) //debug statement
	switch r.Method {
	case http.MethodGet:
		apiCfg.handlerGetUser(w, r)
	case http.MethodPost:
		apiCfg.handlerCreateUser(w, r)
	case http.MethodPut:
		apiCfg.handlerUpdateUser(w, r)
	case http.MethodDelete:
		apiCfg.handlerDeleteUser(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}

type parameters struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      string `json:"age"`
}
