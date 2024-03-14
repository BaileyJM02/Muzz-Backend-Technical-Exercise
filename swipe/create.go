package swipe

import "github.com/baileyjm02/muzz-backend-technical-exercise/context"

// CreateSwipe creates a new swipe
func CreateSwipe(userID int, targetID int, preference string) (Swipe, error) {
	if preference != Like && preference != Dislike {
		return Swipe{}, ErrInvalidPreference
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
