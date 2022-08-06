package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BevansMath/SocialServer/internal/database"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", testHandler)
	database.NewClient("github.com/BevansMath/SocialServer/internal/database")
	const addr = "localhost:8080"
	serv := http.Server{
		Addr:         addr,
		Handler:      m,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Println("server started on", addr)
	err := serv.ListenAndServe()
	log.Fatal(err)

}
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}

func testErrHandler(w http.ResponseWriter, r *http.Request) { // RespondWithErr test handler
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

type errorBody struct {
	Error string `json:"error"` // function takes err, create a new errorBody, then calls respondWithJSON function
}

func respondWithError(w http.ResponseWriter, code int, err error) { // API error handling
	if err == nil {
		log.Println("don't call respondWithError with a nil err")
		return
	}
	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) { //
	w.Header().Set("Content-Type", "applicaton/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	if payload != nil {
		response, err := json.Marshal(payload)
		if err != nil {
			log.Println("error marshalling", err)
			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})
			w.Write(response)
			return
		}
		w.WriteHeader(code)
		w.Write(response)
	}
}
