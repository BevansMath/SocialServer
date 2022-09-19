package main

import (
	"net/http"
)

func (apiCfg apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}
}
