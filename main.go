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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Diretcor *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, "error getMovies", http.StatusInternalServerError)
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, "error deleteMovie", http.StatusInternalServerError)
	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				http.Error(w, "error getMovie", http.StatusInternalServerError)
			}
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "error crateMovie", http.StatusInternalServerError)
	}
	movie.ID = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, "error crateMovie", http.StatusInternalServerError)
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			err := json.NewDecoder(r.Body).Decode(&movie)
			if err != nil {
				http.Error(w, "error updateMovie", http.StatusInternalServerError)
			}
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "First Movie", Diretcor: &Director{Firstname: "Jojo", Lastname: "Bonito"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438789", Title: "Second Movie", Diretcor: &Director{Firstname: "Hoho", Lastname: "Gonito"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
