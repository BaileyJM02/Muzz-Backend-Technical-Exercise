package user

import (
	"github.com/baileyjm02/muzz-backend-technical-exercise/context"
)

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
func GetUnswipedUsers(userID int, filters DiscoverFilters) ([]UserWithDistance, error) {
	ctx := context.GetContext()

	// Get call user for their location
	user, err := GetByID(userID)
	if err != nil {
		return []UserWithDistance{}, err
	}

	users := []UserWithDistance{}
	baseQuery := ctx.DB.Instance.
		Select("id", "name", "gender", "age", "location_longitude", "location_latitude", "location_name").
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

	if user.Location.Latitude != 0 && user.Location.Longitude != 0 {
		// Slightly modified formula from here: https://www.plumislandmedia.net/mysql/haversine-mysql-nearest-loc/
		// The article says that this formula is relatively slow, but it's good enough and clearer than some of the other options.
		baseQuery = baseQuery.Select(`*, 111.045* DEGREES(ACOS(LEAST(1.0, COS(RADIANS(?))
										 * COS(RADIANS(location_latitude))
										 * COS(RADIANS(?) - RADIANS(location_longitude))
										 + SIN(RADIANS(?))
										 * SIN(RADIANS(location_latitude))))) AS distance_in_km`, user.Location.Latitude, user.Location.Longitude, user.Location.Latitude)

		// Order by our newly created column. This also relates to the field in `UserWithDistance`.
		baseQuery = baseQuery.Order("distance_in_km")
	}

	rtx := baseQuery.Find(&users)
	if rtx.Error != nil {
		return []UserWithDistance{}, rtx.Error
	}

	return users, nil
}
