package user

import (
	"fmt"

	"github.com/baileyjm02/muzz-backend-technical-exercise/context"
)

// Location is a struct that represents a high-level location
type Location struct {
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// User is a struct that represents a high-level user
type User struct {
	ID       int      `json:"id" gorm:"primaryKey"`
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty" gorm:"unique"`
	Gender   string   `json:"gender"`
	Age      int      `json:"age"`
	Password string   `json:"password,omitempty"`
	Location Location `json:"location" gorm:"embedded;embeddedPrefix:location_"`
}

// UserWithDistance is a struct that adds a distance field to the User struct
type UserWithDistance struct {
	User           `json:",inline"`
	DistanceFromMe float64 `json:"distance_from_me,omitempty" gorm:"column:distance_in_km"`
}

// TableName is used to define the table name for the User struct in the database
// We want to continue using the 'users' table, as this struct just declares our calculated
// distance column.
func (UserWithDistance) TableName() string {
	return "users"
}

// init is called when the package is imported, and is used to migrate the database to the latest schema
func init() { AutoMigrate() }

// AutoMigrate ensures the database is migrated to the latest schema
func AutoMigrate() {
	ctx := context.GetContext()
	fmt.Println("[USER] [MIGRATION] Starting User Migration...")

	err := ctx.DB.Instance.AutoMigrate(&User{})

	if err != nil {
		fmt.Println("[USER] [MIGRATION] Error Migrating User: ", err)
		panic(err)
	}

	fmt.Println("[USER] [MIGRATION] User Migration Complete!")
}
