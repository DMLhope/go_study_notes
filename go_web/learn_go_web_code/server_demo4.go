package main

import (
	"fmt"
	"net/http"
)

func main5() {
	server := http.Server{
		Addr: "localhost:9000",
	}

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// r.ParseForm()
		// fmt.Fprintln(w, r.Form)
		// fmt.Fprintln(w, r.PostForm)
		r.ParseMultipartForm(1024)
		fmt.Fprintln(w, r.MultipartForm)

	})

	server.ListenAndServe()
}
