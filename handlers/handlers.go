package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID, _ := strconv.ParseInt(vars["chat_id"], 10, 64)
	fmt.Fprintf(w, "You are in %d chat", chatID)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Someone have connected")
	fmt.Fprint(w, "Hello on my chat")
}
