package user

import "github.com/baileyjm02/muzz-backend-technical-exercise/context"

// GetByEmail returns a user by their email
func GetByEmail(email string) (User, error) {
	ctx := context.GetContext()
	user := User{}
	rtx := ctx.DB.Instance.Where("email = ?", email).First(&user)
	if rtx.Error != nil {
		return User{}, rtx.Error
	}

	return user, nil
}

// GetByID returns a user by their ID
func GetByID(id int) (User, error) {
	ctx := context.GetContext()
	user := User{}
	rtx := ctx.DB.Instance.Where("id = ?", id).First(&user)
	if rtx.Error != nil {
		return User{}, rtx.Error
	}

	return user, nil
}

// Get all users that the user hasn't swiped on. (Ensuring the user is not included in the results.)
func GetUnswipedUsers(userID int) ([]User, error) {
	ctx := context.GetContext()
	users := []User{}
	rtx := ctx.DB.Instance.Raw(`
		SELECT id,name,gender,age FROM users
		WHERE id != ?
		AND id NOT IN (
			SELECT target_id FROM swipes WHERE user_id = ?
		)
	`, userID, userID).Scan(&users)
	if rtx.Error != nil {
		return []User{}, rtx.Error
	}

	return users, nil
}
