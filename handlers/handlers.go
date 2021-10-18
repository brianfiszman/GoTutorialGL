package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var response Response = Response{
		Status:  http.StatusOK,
		Message: "Hola Mundo",
	}

	json.NewEncoder(w).Encode(response)
}

func HandleRequests() {
	var port string = os.Getenv("HTTP_PORT")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homePage)
	log.Println("Servirn the app in the port:", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println(err)
	}

}
