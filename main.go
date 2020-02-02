package main

import (
	"fmt"
	c "go-rest-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/api/books", c.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", c.GetBook).Methods("GET")
	r.HandleFunc("/api/books", c.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", c.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", c.DeleteBook).Methods("DELETE")

	fmt.Println("server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
