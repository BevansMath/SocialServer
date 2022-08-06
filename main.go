package main

import (
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

func respondWithJSON(w http.ResponseWriter, code int, payload interface{})
response, err := json.Marshal(payload)
//deal with err here
w.Write(dat)