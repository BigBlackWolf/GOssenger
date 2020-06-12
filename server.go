package main

import (
	"fmt"
	"net/http"

	"GOssenger/chat"
	"GOssenger/dashboard"

	"github.com/gorilla/mux"
)

const (
	siteHost = "127.0.0.1"
	sitePort = 8080
)

func main() {
	address := fmt.Sprintf("%s:%d", siteHost, sitePort)
	fmt.Printf("Starting server on port %s", address)
	createServer(":8080")
}

func createServer(address string) {
	router := mux.NewRouter()
	router.HandleFunc("/{chat_id:[0-9]+}", chat.ChatHandler)
	router.HandleFunc("/", chat.IndexHandler)
	http.Handle("/", router)
	dashboard.CreateRouter(router)

	http.ListenAndServe(address, nil)
}

func Hello() string {
	return "Just for test"
}
