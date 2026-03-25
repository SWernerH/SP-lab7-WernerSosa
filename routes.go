package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// students
	mux.HandleFunc("GET /students", app.listStudents)
	mux.HandleFunc("GET /students/{id}", app.getStudent)
	mux.HandleFunc("POST /students", app.createStudent)
	mux.HandleFunc("PUT /students/{id}", app.updateStudent)
	mux.HandleFunc("DELETE /students/{id}", app.deleteStudent)

	// courses
	mux.HandleFunc("GET /courses", app.listCourses)
	mux.HandleFunc("POST /courses", app.createCourse)

	return mux
}