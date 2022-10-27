package app

import (
	"game-score/internal/transport"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	handlers *transport.Handlers
}

func NewApp(handlers *transport.Handlers) *App {
	return &App{
		handlers: handlers,
	}
}

func (a *App) Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/score/{id}", a.handlers.GetScoreHandler).Methods("GET")
	r.HandleFunc("/score", a.handlers.SaveScoreHandler).Methods("POST")

	log.Println("Starting server on port 8080")

	return http.ListenAndServe(":8080", r)
}
