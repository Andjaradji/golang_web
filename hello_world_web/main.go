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

	log.Println("Starting web port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
