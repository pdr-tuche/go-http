package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	artcles := Articles{
		Article{Title: "teste titulo", Desc: "teste descricao", Content: "teste "},
	}

	fmt.Println("endpoint dos articles")
	json.NewEncoder(w).Encode(artcles)
}

func testPostAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TO CHAMANDO UM POST")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home endpoint aqui")
}

func handleRequests() {
	rota := mux.NewRouter().StrictSlash(true)

	rota.HandleFunc("/", home)
	rota.HandleFunc("/articles", allArticles).Methods("GET")
	rota.HandleFunc("/articles", testPostAllArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", rota))
}

func main() {
	handleRequests()
}
