package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"test"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = ""
	dbName     = "testing"
	dbSslmode  = "disable"
	siteHost   = "127.0.0.1"
	sitePort   = 8080
)

func main() {
	test.Rest()
	address := fmt.Sprintf("%s:%d", siteHost, sitePort)
	fmt.Printf("Starting server on port %s", address)
	createServer(":8080")

	connector := connectToDb()
	fmt.Println(connector)
}

func createServer(address string) {
	router := mux.NewRouter()
	router.HandleFunc("/{chat_id:[0-9]+}", chatHandler)
	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)

	http.ListenAndServe(address, nil)
}

func connectToDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+
		"password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSslmode)

	connector, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return connector
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatID, _ := strconv.ParseInt(vars["chat_id"], 10, 64)
	fmt.Fprintf(w, "You are in %d chat", chatID)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Someone have connected")
	fmt.Fprint(w, "Hello on my chat")
}

type user struct {
	ID       int64
	Username string
	Email    string
}
