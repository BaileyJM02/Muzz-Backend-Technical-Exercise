package user

// User is a struct that represents a high-level user
type User struct {
	ID     int `json:"id"`
	Name   string
	Email  string
	Gender string
	Age    int
}
