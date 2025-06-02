package handler

import (
	"gogame/service"
	"net/http"
	"text/template"
)

func IndexHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "index", nil)
	}
}

func GetGameGridHandler(gs service.GameService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		grid := gs.GetGrid()
		w.Header().Add("content-type", "text/html")
		w.Write([]byte(grid))
	}
}