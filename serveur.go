package main

import (
	"html/template"
	"log"
	"net/http"
)

type student struct {
	name  string
}

type Information struct {
	Lastname  string
	Firstname string
	Age       int
}

func home(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./index.html", "./template/header.html", "./template/footer.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, nil)
}

func Infos (w http.ResponseWriter, r *http.Request, infos *[]Information) {
	template, err := template.ParseFiles("./page/infos.html", "./template/header.html", "./template/footer.html", "./template/information.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func main() {
	var tabInfos []Information
	tabInfos = append(tabInfos, Information{"cozette", "joaquin", 18})
	tabInfos = append(tabInfos, Information{"aich", "gaele", 54})
	tabInfos = append(tabInfos, Information{"cozette", "stephane", 49})
	http.HandleFunc("/", home)
	http.HandleFunc("/infos", func(w http.ResponseWriter, r *http.Request) {
		Infos(w, r, &tabInfos)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)

}
