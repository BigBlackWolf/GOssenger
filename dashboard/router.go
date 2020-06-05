package dashboard

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter(router *mux.Router) *mux.Router {
	sub := router.PathPrefix("/dashboard/").Subrouter()
	sub.HandleFunc("/", indexRouter)

	return router
}

func addTask(w http.ResponseWriter, r *http.Request) {

}

func indexRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello in dashboard")
}

type Task struct {
	id      int32
	title   string
	message string
	userID  int32
}
