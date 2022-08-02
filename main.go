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
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
