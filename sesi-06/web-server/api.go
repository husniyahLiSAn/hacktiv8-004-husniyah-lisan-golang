package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "John", Age: 1, Division: "Art"},
	{ID: 2, Name: "John", Age: 2, Division: "Music"},
	{ID: 3, Name: "John", Age: 3, Division: "Financial"},
}

var PORT = ":8282"

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/employees", getEmployees)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!"
	fmt.Fprint(w, msg)
}

/* func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
} */

func getEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}
