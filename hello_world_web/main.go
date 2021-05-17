package main

import (
	"hello_world_web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloWorldHandler)
	mux.HandleFunc("/about", handler.AboutHandler)
	mux.HandleFunc("/profile", handler.ProfileHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form",handler.Form)
	mux.HandleFunc("/process",handler.Process)


	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/",http.StripPrefix("/static",fileServer))

	log.Println("Starting web port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
