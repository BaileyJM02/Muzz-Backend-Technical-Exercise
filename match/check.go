package match

import (
	"github.com/baileyjm02/muzz-backend-technical-exercise/swipe"
)

// CheckMatch checks if there are two swipes that match each other
func CheckMatch(userOneID int, userTwoID int) (bool, error) {
	swipeOne, err := swipe.GetSwipesByUserIDAndTargetID(userOneID, userTwoID)
	if err != nil {
		return false, nil
	}

	swipeTwo, err := swipe.GetSwipesByUserIDAndTargetID(userTwoID, userOneID)
	if err != nil {
		return false, nil
	}

	if swipeOne.Preference == swipe.Like && swipeTwo.Preference == swipe.Like {
		return true, nil
	}

	return false, nil
}

