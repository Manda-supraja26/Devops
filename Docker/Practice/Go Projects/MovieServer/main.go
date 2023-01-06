package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"/github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var Movies []Movie

func main() {
	// fmt.Printf("Hello world")
	r := mux.NewRouter()
	Movies = append(Movies, Movie{ID: "1", Isbn: "4567a", Title: "Movie One", Director: &Director{FirstName: "Jhon", LastName: "Bob"}})
	Movies = append(Movies, Movie{ID: "2", Isbn: "9843b", Title: "Movie two", Director: &Director{FirstName: "Mathew", LastName: "clara"}})
	r.HandleFunc("/Movies", getMovies).Methods("GET")
	r.HandleFunc("/Movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/Movies", createMovie).Methods("POST")
	r.HandleFunc("/Movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/Movies/{id}", deleteMovie).Methods("DELECTE")
	http.Handle("/", r)

	fmt.Println("Server Listening on 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}

func getMovies(w http.ResponseWriter, r http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)

}

func getMovie(w http.ResponseWriter, r http.Response) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r http.Response) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Movies {
		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(Movies)
}

func createMovie(w http.ResponseWriter, r http.Response) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(Movies)

}

func updateMovie(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Movies {
		if item.ID == params["id"] {
			Movies = append(Movies[:index], Movies[index+1:]...)
		}
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(Movies)

}
