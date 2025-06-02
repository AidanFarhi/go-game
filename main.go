package main

import (
	"gogame/handler"
	"gogame/service"
	"net/http"
	"text/template"
)

func main() {
	t := template.Must(template.ParseGlob("web/html/*.html"))
	fs := http.FileServer(http.Dir("web"))
	m := http.NewServeMux()
	gs := service.NewGameService(20, 20)

	m.Handle("/web/", http.StripPrefix("/web/", fs))
	m.HandleFunc("/", handler.IndexHandler(t))
	m.HandleFunc("GET /gamegrid", handler.GetGameGridHandler(gs))

	s := http.Server{
		Addr: ":7889",
		Handler: m,
	}

	s.ListenAndServe()
}
