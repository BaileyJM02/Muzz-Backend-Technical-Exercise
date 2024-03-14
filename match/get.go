package match

import "github.com/baileyjm02/muzz-backend-technical-exercise/context"

// GetMatchesByID returns a match by its ID.
func GetMatchByID(id int) (Match, error) {
	ctx := context.GetContext()
	match := Match{}
	rtx := ctx.DB.Instance.Where("id = ?", id).First(&match)
	if rtx.Error != nil {
		return Match{}, rtx.Error
	}

	return match, nil
}

// GetMatchesByUserID returns all matches for a user.
func GetMatchesByUserID(userID int) ([]Match, error) {
	ctx := context.GetContext()
	matches := []Match{}
	rtx := ctx.DB.Instance.Where("user_one_id = ? OR user_two_id = ?", userID, userID).Find(&matches)
	if rtx.Error != nil {
		return []Match{}, rtx.Error
	}

	return matches, nil
}
