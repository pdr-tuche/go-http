package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	artcles := Articles{
		Article{Title: "teste titulo", Desc: "teste descricao", Content: "teste conteudo"},
	}

	fmt.Println("endpoint dos articles")
	json.NewEncoder(w).Encode(artcles)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home endpoint aqui")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
