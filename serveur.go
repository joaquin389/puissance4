package main

import (
    "html/template"
    "log"
    "net/http"
)

type GameInfo struct {
    Player1    string
    Player2    string
    Difficulty string
}

func home(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles(
        "./index.html",
        "./template/header.html",
        "./template/footer.html",
    )
    if err != nil {
        log.Fatal(err)
    }
    tmpl.Execute(w, nil)
}

func infos(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        player1 := r.FormValue("player1")
        player2 := r.FormValue("player2")
        difficulty := r.FormValue("difficulty")

        data := GameInfo{
            Player1:    player1,
            Player2:    player2,
            Difficulty: difficulty,
        }

        tmpl, err := template.ParseFiles(
            "./page/infos.html",
            "./template/header.html",
            "./template/footer.html",
        )
        if err != nil {
            log.Fatal(err)
        }
        tmpl.Execute(w, data)
    } else {
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/infos", infos)

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("Serveur lancé sur http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
