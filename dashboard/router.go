package dashboard

import (
	"GOssenger/dashboard/handlers"

	"github.com/gorilla/mux"
)

func CreateRouter(router *mux.Router) *mux.Router {
	sub := router.PathPrefix("/dashboard").Subrouter()
	sub.HandleFunc("", handlers.IndexHandler).Methods("GET", "OPTIONS")
	sub.HandleFunc("/", handlers.IndexHandler).Methods("GET", "OPTIONS")
	sub.HandleFunc("", handlers.AddTaskHandler).Methods("POST", "OPTIONS")
	sub.HandleFunc("/", handlers.AddTaskHandler).Methods("POST", "OPTIONS")
	sub.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST", "OPTIONS")
	sub.HandleFunc("/register", handlers.RegisterHandler).Methods("GET", "POST", "OPTIONS")
	sub.HandleFunc("/{title}", handlers.TaskHandler).Methods("GET", "OPTIONS")
	sub.HandleFunc("/{title}", handlers.DeleteTaskHandler).Methods("DELETE", "OPTIONS")
	return router
}
