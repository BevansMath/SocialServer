package main 

package database

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", testHandler)

	const addr = "localhost:8080"
	serv := http.Server{
		Addr:    addr,
		Handler: m,
		//TLSConfig:         &tls.Config{},
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		//IdleTimeout:       0,
		//MaxHeaderBytes:    0,
		//TLSNextProto:      map[string]func(*http.Server, *tls.Conn, http.Handler){},
		//ConnState: func(net.Conn, http.ConnState) {
		//},
		//ErrorLog: &log.Logger{},
		//BaseContext: func(net.Listener) context.Context {
		//},
		//ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		//},
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
