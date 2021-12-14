package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	handlerEmployee "sample-api/handlers/employee"
	handlerStudent "sample-api/handlers/student"
	serviceEmployee "sample-api/services/employee"
	serviceStudent "sample-api/services/student"
	storeEmployee "sample-api/stores/employee"
	storeStudent "sample-api/stores/student"
)

func main() {
	const PORT = 8000

	var db *sql.DB //fixme

	r := mux.NewRouter()

	studentStore := storeStudent.New(db)
	studentService := serviceStudent.New(studentStore)
	studentHandler := handlerStudent.New(studentService)

	employeeStore := storeEmployee.New(db)
	employeeService := serviceEmployee.New(employeeStore)
	employeeHandler := handlerEmployee.New(employeeService)

	// register routes for students
	r.HandleFunc("/students", studentHandler.Create).Methods(http.MethodPost)        // create a new student record
	r.HandleFunc("/students", studentHandler.GetAll).Methods(http.MethodGet)         // get all student records
	r.HandleFunc("/students/{id}", studentHandler.Get).Methods(http.MethodGet)       // get student by ID
	r.HandleFunc("/students/{id}", studentHandler.Update).Methods(http.MethodPut)    // update student by ID
	r.HandleFunc("/students/{id}", studentHandler.Delete).Methods(http.MethodDelete) // delete student by ID

	// register routes for employees
	r.HandleFunc("/employees", employeeHandler.Create).Methods(http.MethodPost)        // create a new employee record
	r.HandleFunc("/employees", employeeHandler.GetAll).Methods(http.MethodGet)         // get all employee records
	r.HandleFunc("/employees/{id}", employeeHandler.Get).Methods(http.MethodGet)       // get employee by ID
	r.HandleFunc("/employees/{id}", employeeHandler.Update).Methods(http.MethodPut)    // update employee by ID
	r.HandleFunc("/employees/{id}", employeeHandler.Delete).Methods(http.MethodDelete) // delete employee by ID

	srv := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", PORT),
		Handler: r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error in starting server: %v\n", err)
	}
}
