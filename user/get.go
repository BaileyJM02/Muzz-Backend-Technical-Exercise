package user

import "github.com/baileyjm02/muzz-backend-technical-exercise/context"

func GetByEmail(email string) (User, error) {
	ctx := context.GetContext()
	user := User{}
	rtx := ctx.DB.Instance.Where("email = ?", email).First(&user)
	if rtx.Error != nil {
		return User{}, rtx.Error
	}

	return user, nil
}
