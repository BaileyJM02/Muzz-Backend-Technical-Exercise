package match

import (
	"fmt"

	"github.com/baileyjm02/muzz-backend-technical-exercise/context"
)

// Match represents a match between two users when they have both liked each other
type Match struct {
	ID        int `json:"id" gorm:"primaryKey"`
	UserOneID int `json:"user_one_id"`
	UserTwoID int `json:"user_two_id"`
}

// AutoMigrate ensures the database is migrated to the latest schema
func AutoMigrate() {
	ctx := context.GetContext()
	fmt.Println("[MATCH] [MIGRATION] Starting Match Migration...")

	err := ctx.DB.Instance.AutoMigrate(&Match{})

	if err != nil {
		fmt.Println("[MATCH] [MIGRATION] Error Migrating Match: ", err)
		panic(err)
	}

	fmt.Println("[MATCH] [MIGRATION] Match Migration Complete!")
}
