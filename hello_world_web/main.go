package main

import (
	"log"
	"net/http"
)

func main () {
	mux := http.NewServeMux()

	aboutHandler := func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("About Page"))

	}

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloWorldHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Page"))

	})
	log.Println("Starting web port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"{
		w.Write([]byte("Not Found Aja"))
		return
	}
	w.Write([]byte("Welcome To Home"))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World made for Golang Web"))
}
