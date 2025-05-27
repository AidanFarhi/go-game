package main

import (
	"gogame/handler"
	"net/http"
	"text/template"
)

func main() {
	t := template.Must(template.ParseGlob("web/html/*.html"))
	fs := http.FileServer(http.Dir("web"))
	m := http.NewServeMux()

	m.Handle("/web/", http.StripPrefix("/web/", fs))
	m.HandleFunc("/", handler.IndexHandler(t))

	s := http.Server{
		Addr: ":7889",
		Handler: m,
	}

	s.ListenAndServe()
}
