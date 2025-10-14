package main

import (
    "html/template"
    "net/http"
    "strconv"
)

type Game struct {
    Grid     [][]int
    Turn     int
    Winner   int
    Player1  string
    Player2  string
    Rows     int
    Cols     int
    Difficulty string
}

var game *Game
var tmplPlay = template.Must(template.ParseFiles("play.html"))
var tmplHome = template.Must(template.ParseFiles("home.html"))

func NewGame(rows, cols int, p1, p2, difficulty string) *Game {
    grid := make([][]int, rows)
    for i := range grid {
        grid[i] = make([]int, cols)
    }
    return &Game{
        Grid: grid, Turn: 1, Winner: 0,
        Player1: p1, Player2: p2,
        Rows: rows, Cols: cols,
        Difficulty: difficulty,
    }
}

func (g *Game) Play(col int) {
    if g.Winner != 0 {
        return
    }
    for row := g.Rows - 1; row >= 0; row-- {
        if g.Grid[row][col] == 0 {
            g.Grid[row][col] = g.Turn
            if g.CheckVictory(g.Turn) {
                g.Winner = g.Turn
            } else {
                g.Turn = 3 - g.Turn
            }
            break
        }
    }
}

func (g *Game) CheckVictory(p int) bool {
    R, C := g.Rows, g.Cols
    G := g.Grid

    for r := 0; r < R; r++ {
        for c := 0; c <= C-4; c++ {
            if G[r][c] == p && G[r][c+1] == p && G[r][c+2] == p && G[r][c+3] == p {
                return true
            }
        }
    }
    for c := 0; c < C; c++ {
        for r := 0; r <= R-4; r++ {
            if G[r][c] == p && G[r+1][c] == p && G[r+2][c] == p && G[r+3][c] == p {
                return true
            }
        }
    }
    for r := 0; r <= R-4; r++ {
        for c := 0; c <= C-4; c++ {
            if G[r][c] == p && G[r+1][c+1] == p && G[r+2][c+2] == p && G[r+3][c+3] == p {
                return true
            }
        }
    }
    for r := 3; r < R; r++ {
        for c := 0; c <= C-4; c++ {
            if G[r][c] == p && G[r-1][c+1] == p && G[r-2][c+2] == p && G[r-3][c+3] == p {
                return true
            }
        }
    }
    return false
}

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            p1 := r.FormValue("player1")
            p2 := r.FormValue("player2")
            diff := r.FormValue("difficulty")
            var rows, cols int
            switch diff {
            case "facile":
                rows, cols = 6, 7
            case "moyen":
                rows, cols = 6, 9
            case "difficile":
                rows, cols = 7, 8
            default:
                rows, cols = 6, 7
            }
            game = NewGame(rows, cols, p1, p2, diff)
            http.Redirect(w, r, "/play", http.StatusSeeOther)
            return
        }
        tmplHome.Execute(w, nil)
    })

    http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            colStr := r.FormValue("col")
            col, _ := strconv.Atoi(colStr)
            game.Play(col)
            http.Redirect(w, r, "/play", http.StatusSeeOther)
            return
        }
        tmplPlay.Execute(w, game)
    })

    http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/", http.StatusSeeOther)
    })

    http.ListenAndServe(":8080", nil)
}
