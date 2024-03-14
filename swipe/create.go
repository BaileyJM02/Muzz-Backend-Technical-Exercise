package swipe

import (
	"fmt"

	"github.com/baileyjm02/muzz-backend-technical-exercise/context"
	"github.com/baileyjm02/muzz-backend-technical-exercise/user"
)

var (
	// ErrSwipeExists is returned when a swipe already exists
	ErrSwipeExists = fmt.Errorf("swipe already exists")
)

// CreateSwipe creates a new swipe
func CreateSwipe(userID int, targetID int, preference string) (Swipe, error) {
	if preference != Like && preference != Dislike {
		return Swipe{}, ErrInvalidPreference
	}

	// Check target exists
	_, err := user.GetByID(targetID)
	if err != nil {
		return Swipe{}, err
	}

	// Check swipe doesn't already exist
	_, err = GetSwipesByUserIDAndTargetID(userID, targetID)
	if err == nil {
		return Swipe{}, ErrSwipeExists
	}

	ctx := context.GetContext()
	swipe := Swipe{
		UserID:     userID,
		TargetID:   targetID,
		Preference: preference,
	}
	rtx := ctx.DB.Instance.Create(&swipe)
	if rtx.Error != nil {
		return Swipe{}, rtx.Error
	}

	return swipe, nil
}
