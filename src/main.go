package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	err := InitializeDB("postgres://username:password@localhost/studentsdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer CloseDB()

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/students", CreateStudent).Methods("POST")
	r.HandleFunc("/students", GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id}", GetStudentByID).Methods("GET")
	r.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
