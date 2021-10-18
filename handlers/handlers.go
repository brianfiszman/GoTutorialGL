package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Movie struct {
	Title       string `json:"title"`
	ReleaseDate int    `json:"ReleaseDate"`
	Platform    string `json:"Platform"`
}

var movies = []Movie{
	{Title: "Back to the Future", ReleaseDate: 1985, Platform: "Netflix"},
	{Title: "Joker", ReleaseDate: 2019, Platform: "Prime Video"},
	{Title: "Click", ReleaseDate: 2006, Platform: "HBO Max"},
}

func allMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: All Movies Endpoint")
	json.NewEncoder(w).Encode(movies)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage")
}

func HandleRequests() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homePage)
	r.Get("/movies", allMovies)
	http.ListenAndServe(":8001", r)
}
