package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process_1(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(1024)

	// fileHeader := r.MultipartForm.File["uploaded"][0]
	// file, err := fileHeader.Open()

	file, _, err := r.FormFile("uploaded") //返回上传的第一个文件
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main6() {
	server := http.Server{
		Addr: "localhost:9000",
	}
	http.HandleFunc("/process", process_1)
	server.ListenAndServe()
}
