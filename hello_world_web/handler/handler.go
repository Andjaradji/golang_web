package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Write([]byte("Not Found Aja"))
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"),path.Join("views","layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"title":   "Title Golang Web dari Data",
		"content": "Content Golang Web dari Data",
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
		return
	}

}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World made for Golang Web"))
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile Page"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
		return
	}

	data := map[string]int{
		"content": idNum,
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}
