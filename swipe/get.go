package swipe

import "github.com/baileyjm02/muzz-backend-technical-exercise/context"

// GetSwipesByUserID returns all swipes for a user
func GetSwipesByUserID(userID int) ([]Swipe, error) {
	ctx := context.GetContext()
	swipes := []Swipe{}
	rtx := ctx.DB.Instance.Where("user_id = ?", userID).Find(&swipes)
	if rtx.Error != nil {
		return []Swipe{}, rtx.Error
	}

	return swipes, nil
}

// GetSwipesByTargetID returns all swipes for a target
func GetSwipesByTargetID(targetID int) ([]Swipe, error) {
	ctx := context.GetContext()
	swipes := []Swipe{}
	rtx := ctx.DB.Instance.Where("target_id = ?", targetID).Find(&swipes)
	if rtx.Error != nil {
		return []Swipe{}, rtx.Error
	}

	return swipes, nil
}

// GetSwipesByUserIDAndTargetID returns swipe for a user and target if it exists
func GetSwipesByUserIDAndTargetID(userID int, targetID int) (Swipe, error) {
	ctx := context.GetContext()
	swipes := Swipe{}
	rtx := ctx.DB.Instance.Where("user_id = ? AND target_id = ?", userID, targetID).First(&swipes)
	if rtx.Error != nil {
		return Swipe{}, rtx.Error
	}

	return swipes, nil
}
