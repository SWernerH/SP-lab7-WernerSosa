package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, values := range headers {
		for _, v := range values {
			w.Header().Add(key, v)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return errors.New("empty body")
		}
		if strings.Contains(err.Error(), "unknown field") {
			return errors.New("unknown field in JSON")
		}
		return err
	}
	return nil
}

// errors

func (app *application) serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	app.writeJSON(w, 500, envelope{"error": "server error"}, nil)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.writeJSON(w, 404, envelope{"error": "not found"}, nil)
}

func (app *application) badRequest(w http.ResponseWriter, msg string) {
	app.writeJSON(w, 400, envelope{"error": msg}, nil)
}