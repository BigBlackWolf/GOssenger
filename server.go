package main

import (
	"fmt"
	"net/http"

	"GOssenger/chat"
	"GOssenger/dashboard"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	siteHost = "127.0.0.1"
	sitePort = 8080
)

func main() {
	address := fmt.Sprintf("%s:%d", siteHost, sitePort)
	fmt.Println(fmt.Sprintf("Starting server on port %s", address))
	createServer(":8080")
}

func createServer(address string) {
	router := mux.NewRouter()
	router.HandleFunc("/{chat_id:[0-9]+}", chat.ChatHandler)
	router.HandleFunc("/", chat.IndexHandler)
	http.Handle("/", router)
	dashboard.CreateRouter(router)

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	http.ListenAndServe(address, corsWrapper.Handler(router))
}

func Hello() string {
	return "Just for test"
}
