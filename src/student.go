package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// CreateStudent creates a new student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO students (name, age, email) VALUES ($1, $2, $3) RETURNING id"
	err = db.QueryRow(query, student.Name, student.Age, student.Email).Scan(&student.ID)
	if err != nil {
		http.Error(w, "Could not insert student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// GetAllStudents retrieves all students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, age, email FROM students")
	if err != nil {
		http.Error(w, "Could not retrieve students", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		if err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Email); err != nil {
			http.Error(w, "Error scanning student", http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	json.NewEncoder(w).Encode(students)
}

// GetStudentByID retrieves a student by ID
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := "SELECT id, name, age, email FROM students WHERE id = $1"
	row := db.QueryRow(query, id)

	var student Student
	err := row.Scan(&student.ID, &student.Name, &student.Age, &student.Email)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

// UpdateStudent updates an existing student
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "UPDATE students SET name = $1, age = $2, email = $3 WHERE id = $4"
	_, err = db.Exec(query, student.Name, student.Age, student.Email, id)
	if err != nil {
		http.Error(w, "Could not update student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteStudent deletes a student by ID
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := "DELETE FROM students WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, "Could not delete student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
