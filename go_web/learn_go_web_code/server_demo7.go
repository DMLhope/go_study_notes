package main

import (
	"net/http"
	"text/template"
)

func main8() {
	server := http.Server{
		Addr: "localhost:9000",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "hello World!")
}
