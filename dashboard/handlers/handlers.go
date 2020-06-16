package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"GOssenger/dashboard/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbHost     = "localhost"
	dbPort     = 27017
	dbUser     = "mongodb"
	dbPassword = ""
	dbName     = "testing"
	dbSslmode  = "disable"
)

var collection *mongo.Collection

func init() {
	// Connect to db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	collection = client.Database("Dashboard").Collection("myCollection")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	options := options.Find().SetSort(bson.D{{"_id", -1}})
	curr, err := collection.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		log.Fatal(nil)
	}
	var results []models.Task
	for curr.Next(context.TODO()) {
		var result models.Task
		e := curr.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}

	// sort.Sort(models.ByID(results))
	curr.Close(context.TODO())
	json.NewEncoder(w).Encode(results)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, OPTIONS, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	task_id, _ := strconv.Atoi(vars["task_id"])
	var result models.Task
	filter := bson.D{{"id", task_id}}
	collection.FindOne(context.TODO(), filter).Decode(&result)
	json.NewEncoder(w).Encode(result)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, OPTIONS, GET")
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var res models.Task
	err := decoder.Decode(&res)
	if err != nil {
		log.Fatal(err)
	}

	insertResult, err := collection.InsertOne(context.TODO(), res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Record ", insertResult.InsertedID)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["task_id"])

	filter := bson.M{"id": id}
	d, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document ", d.DeletedCount)
}
