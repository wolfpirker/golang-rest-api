package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wolfpirker/golang-rest-api/routes"
)

func main() {
	var r = mux.NewRouter()
	const port string = ":8000"
	r.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
		// Note: fprintln to show message to the client
	})
	r.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	r.HandleFunc("/posts", routes.AddPost).Methods("POST")
	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, r))
}
