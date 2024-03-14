package user

import (
	"fmt"
	"strings"

	"github.com/baileyjm02/muzz-backend-technical-exercise/context"
	"github.com/brianvoe/gofakeit/v7"
)

func create(newUser User) User {
	ctx := context.GetContext()

	ctx.DB.Instance.Create(&newUser)

	// Get user from database
	var u User
	ctx.DB.Instance.Where("email = ?", newUser.Email).First(&u)

	return u
}

func CreateRandom() User {
	fake := gofakeit.New(0)
	fakePerson := fake.Person()

	name := fmt.Sprintf("%v %v", fakePerson.FirstName, fakePerson.LastName)
	email := fmt.Sprintf("%v@%v.muzz", fakePerson.FirstName, fakePerson.LastName)

	user := User{
		Name:     name,
		Email:    strings.ToLower(email),
		Gender:   fake.Gender(),
		Age:      fake.Number(18, 40),
		Password: fake.Password(true, true, true, false, false, 10),
	}

	return create(user)
}
