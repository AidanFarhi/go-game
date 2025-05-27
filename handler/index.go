package handler

import (
	"net/http"
	"text/template"
)

func IndexHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index", nil)
	}
}