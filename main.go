package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BevansMath/SocialServer/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

func main() {
	m := http.NewServeMux()

	const dbPath = "db.json" // Ensure the database exists
	dbClient := database.NewClient(dbPath)

	err := dbClient.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{ // Configure the endpoints between the client and the database
		dbClient: dbClient,
	}

	m.HandleFunc("/users", apiCfg.endpointUsersHandler)
	m.HandleFunc("/users/", apiCfg.endpointUsersHandler)
	m.HandleFunc("/err", testErrHandler)
	m.HandleFunc("/test", testHandler)
	m.HandleFunc("/posts", apiCfg.endpointPostHandler)
	m.HandleFunc("/posts/", apiCfg.endpointPostHandler)

	const addr = "localhost:8080"
	serv := http.Server{
		Addr:         addr,
		Handler:      m,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Println("server started on", addr) // This will block forever,
	err = serv.ListenAndServe()            // Until the server has an unrecoverable error
	log.Fatal(err)

}

func testHandler(w http.ResponseWriter, r *http.Request) { // Tests the 200 OK! response
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func testErrHandler(w http.ResponseWriter, r *http.Request) { // Tests 500 server error unrecoverable
	respondWithError(w, 500, errors.New("server encountered a fatal error"))
}

type errorBody struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("dont call respondWithError with a nil err!")
		return
	}
	log.Println(err)
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) { // Endpoint header payload information
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if payload != nil {
		response, err := json.Marshal(payload) // Converts var type payload into json string for response header
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
