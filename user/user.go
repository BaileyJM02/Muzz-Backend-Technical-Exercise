package user

// User is a struct that represents a high-level user
type User struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
