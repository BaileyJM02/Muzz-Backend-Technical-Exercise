package user

import (
	"fmt"

	"github.com/baileyjm02/muzz-backend-technical-exercise/context"
)

// User is a struct that represents a high-level user
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty" gorm:"unique"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	Password string `json:"password,omitempty"`
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
