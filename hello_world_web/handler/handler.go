package handler

import (
	"hello_world_web/entity"
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

	// data := entity.Product{ID: idNum, Name: "Byson", Price: 12000000, Manufacture: "Yamaha"}
	data := []entity.Product{
		{ID: idNum, Name: "Byson", Price: 12000000, Manufacture: "Yamaha", Stock: 11},
		{ID: idNum, Name: "Tiger", Price: 18000000, Manufacture: "Honda", Stock: 8},
		{ID: idNum, Name: "Scorpio", Price: 10000000, Manufacture: "Yamaha", Stock: 3},
		{ID: idNum, Name: "Bandit", Price: 15000000, Manufacture: "Suzuki", Stock: 2},
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

func PostGet(w http.ResponseWriter, r *http.Request){
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini Adalah Get"))
	case "POST":
		w.Write([]byte("Ini Adalah Post"))
	default:
		http.Error(w, "Error is Happenning. Keep Calm", http.StatusBadRequest)
	}
}

func Form (w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w,"Error is Happening. Keep Calm", http.StatusBadRequest)

}

func Process (w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name" : name,
			"message" : message,
		}

		tmpl,err := template.ParseFiles(path.Join("views","result.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Terjadi Kesalahan pada sistem", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w,"Error is Happening. Keep Calm", http.StatusBadRequest)

}
