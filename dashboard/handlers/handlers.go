package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"GOssenger/dashboard/models"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	result := models.GetAllTasks()
	json.NewEncoder(w).Encode(result)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fields := make(map[string]interface{})
	fields["title"] = vars["title"]
	result := models.GetTask(fields)
	json.NewEncoder(w).Encode(result)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var res models.Task
	err := decoder.Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	models.AddTask(res)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fields := make(map[string]interface{})
	fields["title"] = vars["title"]
	models.DeleteTask(fields)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var data map[string]string
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	username := data["username"]
	password := data["password"]
	user := models.GetUser(username, password)
	json.NewEncoder(w).Encode(user)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	user.Register()
}
