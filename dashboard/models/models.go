package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID      int                `json:"id,omitempty"`
	Title   string             `json:"title,omitempty"`
	Message string             `json:"message,omitempty"`
	UserID  int                `json:"userID,omitempty"`
}

type User struct {
	ID       int
	Username string
	Email    string
}

type ByID []Task

func (a ByID) Len() int { return len(a) }
func (a ByID) Less(i, j int) bool {
	if a[i].ID == a[j].ID {
		return a[i].Title < a[j].Title
	} else {
		return a[i].ID < a[j].ID
	}
}
func (a ByID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
