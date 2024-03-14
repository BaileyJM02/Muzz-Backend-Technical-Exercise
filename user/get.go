package user

import "github.com/baileyjm02/muzz-backend-technical-exercise/context"

// DiscoverFilters that we can use in the GetUnswipedUsers function.
type DiscoverFilters struct {
	MinAge int
	MaxAge int
	Gender string
}

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
func GetUnswipedUsers(userID int, filters DiscoverFilters) ([]User, error) {
	ctx := context.GetContext()
	users := []User{}
	baseQuery := ctx.DB.Instance.
		Select("id", "name", "gender", "age").
		Where("id NOT IN (SELECT target_id FROM swipes WHERE user_id = ?) AND id != ?", userID, userID)

	if filters.MinAge > 0 {
		baseQuery = baseQuery.Where("age >= ?", filters.MinAge)
	}

	if filters.MaxAge > 0 {
		baseQuery = baseQuery.Where("age <= ?", filters.MaxAge)
	}

	if filters.Gender != "" {
		baseQuery = baseQuery.Where("gender = ?", filters.Gender)
	}

	rtx := baseQuery.Find(&users)
	if rtx.Error != nil {
		return []User{}, rtx.Error
	}

	return users, nil
}
