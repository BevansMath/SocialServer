package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerGetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("retrieving post")
	postID := strings.TrimPrefix(r.URL.Path, "/posts/")
	if postID == "" {
		fmt.Println("error, get request made for non-existent postID")
		respondWithError(w, http.StatusBadRequest, errors.New("cannot make request without postID"))
		return
	}
	posts, err := apiCfg.dbClient.GetPosts(postID)
	if err != nil {
		fmt.Println("Bad parameters")
		respondWithError(w, http.StatusBadRequest, err)
		return
	}
	respondWithJSON(w, http.StatusOK, posts)
}
