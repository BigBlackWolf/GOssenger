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
	sub.HandleFunc("/{task_id:[0-9]+}", handlers.TaskHandler).Methods("GET", "OPTIONS")
	sub.HandleFunc("/{task_id:[0-9]+}", handlers.DeleteTaskHandler).Methods("DELETE", "OPTIONS")
	return router
}
