package server

import (
	"../kafka"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"time"
)

// var ServeMux serveMux

func WebServer() {
	r := mux.NewRouter()
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		fmt.Println(vars["videoUrl"], "vars")
		fmt.Println(req.URL, "request")
		fmt.Println(req.URL.Query()["a"], "request")
		go kafka.Producer(vars["videoUrl"])
		io.WriteString(w, "Hello, world!!!\n")
	}

	r.HandleFunc("/hello/{videoUrl}", helloHandler)
	r.HandleFunc("/hello", helloHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// Good practice: enforce timeouts for servers you create!
	}
	log.Fatal(srv.ListenAndServe())
}
