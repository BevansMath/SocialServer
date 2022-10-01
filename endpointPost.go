package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endpointPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiCfg.handlerGetPost(w, r)
	case http.MethodPost:
		apiCfg.handlerCreatePost(w, r)
	case http.MethodDelete:
		apiCfg.handlerDeletePost(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}
