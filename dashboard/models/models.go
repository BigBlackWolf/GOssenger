package models

type Task struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
	UserID  int    `json:"userID,omitempty"`
}

type User struct {
	ID       int
	Username string
	Email    string
}
