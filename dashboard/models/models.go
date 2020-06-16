package models

type Task struct {
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
	UserID  int    `json:"userID,omitempty"`
}

type User struct {
	ID       int
	Username string
	Email    string
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
