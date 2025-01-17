package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	firstName string
	lastName  string
}

var movies []Movie

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["Id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}

	}
}

func createmovie(w http.ResponseWriter, r*http.Request){
w.Header().Set("content-type", "application/json")
var movie Movie
json.NewDecoder(r.Body).Decode(&movie)
movie.Id = strconv.Itoa(rand.Intn(10000000))
movies = append(movies,movie )

json.NewEncoder(w).Encode(movie)
}

func updatemovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range movies {
        if item.Id == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            var movie Movie
            _ = json.NewDecoder(r.Body).Decode(&movie)
            movie.Id = params["id"]
            movies = append(movies, movie)
            json.NewEncoder(w).Encode(movie)
            return
        }
    }
    json.NewEncoder(w).Encode(movies)
}

func main() {
	movies = append(movies, Movie{Id: "1", Title: "Interstallar", Director: &Director{firstName: "Christopher", lastName: "Nolan"}})
	movies = append(movies, Movie{Id: "2", Title: "Dune2", Director: &Director{firstName: "Denis", lastName: "Villenuve"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PATCH")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
