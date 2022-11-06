package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `
	<html>
	<head>
	   <title>test</title>
	</head>
	<body>
		<h1>Go web</h1>
	</body>
	</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service,try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("location", "http://www.google.com")
	w.WriteHeader(302)
}

type Post struct {
	User   string
	Theads []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:   "test",
		Theads: []string{"1", "2", "3"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main7() {
	server := http.Server{
		Addr: "localhost:9000",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)

	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
