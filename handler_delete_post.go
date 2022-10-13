package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Delete post from db.json file
func (apiCfg apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	postID := strings.TrimPrefix(r.URL.Path, "/posts/")
	if postID == "" {
		fmt.Println("error, email empty")
		respondWithError(w, http.StatusBadRequest, errors.New("no userEmail provided to handlerDeletePost"))
		return
	}
	err := apiCfg.dbClient.DeletePost(postID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct{}{})
}
