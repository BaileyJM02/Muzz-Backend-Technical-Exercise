package swipe

import (
	"errors"
	"fmt"

	"github.com/baileyjm02/muzz-backend-technical-exercise/context"
)

// Swipe is the idea of 'liking' or 'disliking' another user.
type Swipe struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	UserID     int    `json:"user_id"`
	TargetID   int    `json:"target_id"`
	Preference string `json:"preference" gorm:"type:enum('like', 'dislike')"`
}

// Preference constants
const (
	Like    = "like"
	Dislike = "dislike"
)

// Specific errors we can match against
var ErrInvalidPreference = errors.New("Invalid preference")

// AutoMigrate ensures the database is migrated to the latest schema
func AutoMigrate() {
	ctx := context.GetContext()
	fmt.Println("[SWIPE] [MIGRATION] Starting Swipe Migration...")

	err := ctx.DB.Instance.AutoMigrate(&Swipe{})

	if err != nil {
		fmt.Println("[SWIPE] [MIGRATION] Error Migrating Swipe: ", err)
		panic(err)
	}

	fmt.Println("[SWIPE] [MIGRATION] Swipe Migration Complete!")
}
