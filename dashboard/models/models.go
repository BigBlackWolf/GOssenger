package models

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

var myCollection, userCollection *mongo.Collection

type Task struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Title   string             `json:"title"`
	Message string             `json:"message,omitempty"`
	UserID  int                `json:"userID,omitempty"`
}

func (task *Task) Validate() error {
	a := primitive.NewObjectID()
	fmt.Println(a)
	check := map[string]interface{}{"title": task.Title}
	if GetTask(check) != (Task{}) {
		return errors.New("Not unique")
	}
	if task.Title == "" {
		return errors.New("Empty title")
	}

	task.ID = primitive.NewObjectID()
	fmt.Println("Inserted task, taskId:", task.ID)
	return nil
}

type User struct {
	ID       primitive.ObjectID `json:"_id"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
}

func (user *User) Register() {
	user.ID = primitive.NewObjectID()
	insertResult, err := myCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user created: ", insertResult)
}

func GetUser(username, password string) User {
	var result User
	filter := bson.D{{"username", username}, {"password", password}}
	myCollection.FindOne(context.TODO(), filter).Decode(&result)
	return result
}

type Token struct {
	userId uint
}

func Init() {
	// Connect to db
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	myCollection = client.Database("Dashboard").Collection("myCollection")
	userCollection = client.Database("Dashboard").Collection("user")
}

func GetAllTasks() []Task {
	var results []Task
	options := options.Find().SetSort(bson.D{{"_id", -1}})
	curr, err := myCollection.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		log.Fatal(nil)
	}
	for curr.Next(context.TODO()) {
		var result Task
		e := curr.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}

	// sort.Sort(models.ByID(results))
	curr.Close(context.TODO())
	return results
}

func GetTask(fields map[string]interface{}) Task {
	var result Task
	filter := bson.D{{"title", fields["title"]}}
	myCollection.FindOne(context.TODO(), filter).Decode(&result)
	return result
}

func AddTask(res Task) {
	err := res.Validate()
	if err != nil {
		log.Fatal(err)
	}

	insertResult, err := myCollection.InsertOne(context.TODO(), res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted result", insertResult.InsertedID)
}

func DeleteTask(fields map[string]interface{}) {
	filter := bson.D{{"title", fields["title"]}}
	// for key, value := range fields {
	// 	filter[key] = value
	// }
	d, err := myCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document ", d.DeletedCount)
}

// type ByID []Task

// func (a ByID) Len() int { return len(a) }
// func (a ByID) Less(i, j int) bool {
// 	if a[i].ID == a[j].ID {
// 		return a[i].Title < a[j].Title
// 	} else {
// 		return a[i].ID < a[j].ID
// 	}
// }
// func (a ByID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
