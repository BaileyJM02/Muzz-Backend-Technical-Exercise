package match

import "github.com/baileyjm02/muzz-backend-technical-exercise/context"

// CreateMatch creates a new match between two users
func CreateMatch(userOneID int, userTwoID int) (Match, error) {
	ctx := context.GetContext()
	match := Match{
		UserOneID: userOneID,
		UserTwoID: userTwoID,
	}
	rtx := ctx.DB.Instance.Create(&match)
	if rtx.Error != nil {
		return Match{}, rtx.Error
	}

	return match, nil
}