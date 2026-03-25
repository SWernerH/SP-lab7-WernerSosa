package main

import (
	"context"
	"net/http"
	"time"
)

func (app *application) listCourses(w http.ResponseWriter, r *http.Request) {

	query := `SELECT code, title, credits, enrolled FROM courses`

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	rows, err := app.db.QueryContext(ctx, query)
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var c Course
		rows.Scan(&c.Code, &c.Title, &c.Credits, &c.Enrolled)

		instRows, _ := app.db.QueryContext(ctx,
			`SELECT instructor FROM course_instructors WHERE course_code=$1`, c.Code)

		for instRows.Next() {
			var name string
			instRows.Scan(&name)
			c.Instructors = append(c.Instructors, name)
		}
		instRows.Close()

		courses = append(courses, c)
	}

	app.writeJSON(w, http.StatusOK, envelope{"courses": courses}, nil)
}

func (app *application) createCourse(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Code        string   `json:"code"`
		Title       string   `json:"title"`
		Credits     int      `json:"credits"`
		Instructors []string `json:"instructors"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	_, err = app.db.ExecContext(ctx,
		`INSERT INTO courses (code, title, credits) VALUES ($1,$2,$3)`,
		input.Code, input.Title, input.Credits)

	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, inst := range input.Instructors {
		app.db.ExecContext(ctx,
			`INSERT INTO course_instructors (course_code, instructor) VALUES ($1,$2)`,
			input.Code, inst)
	}

	app.writeJSON(w, http.StatusCreated, envelope{"course": input}, nil)
}