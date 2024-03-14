package user

import (
	"fmt"
	"strings"

	"github.com/baileyjm02/muzz-backend-technical-exercise/context"
	"github.com/baileyjm02/muzz-backend-technical-exercise/utils"
	"github.com/brianvoe/gofakeit/v7"
)

// Create a new user, ensuring to hash the password and return it after creation
func Create(newUser User) (User, error) {
	var err error
	ctx := context.GetContext()

	// Hash the password for the DB
	newUser.Password, err = utils.GenerateHashPassword(newUser.Password)
	if err != nil {
		fmt.Println("error hashing password: ", err)
		return User{}, err
	}

	// Create user in database
	wtx := ctx.DB.Instance.Create(&newUser)
	if wtx.Error != nil {
		fmt.Println("error creating user: ", wtx.Error)
		return User{}, wtx.Error
	}

	// Get user from database
	var createdUser User
	rtx := ctx.DB.Instance.Where("email = ?", newUser.Email).First(&createdUser)
	if rtx.Error != nil {
		fmt.Println("error getting user: ", rtx.Error)
		return User{}, rtx.Error
	}

	return createdUser, nil
}

// CreateRandom creates a user with random data, returning the user as if you had called 'Create'.
func CreateRandom() (User, error) {
	// Initiate the fake data generator
	fake := gofakeit.New(0)
	fakePerson := fake.Person()

	name := fmt.Sprintf("%v %v", fakePerson.FirstName, fakePerson.LastName)
	email := fmt.Sprintf("%v@%v.muzz", fakePerson.FirstName, fakePerson.LastName)

	// Store so we can alter the response as we want to be able to actually login
	plainTextPassword := fake.Password(true, true, true, false, false, 10)

	// Populate user fields
	user := User{
		Name:     name,
		Email:    strings.ToLower(email),
		Gender:   fake.Gender(),
		Age:      fake.Number(18, 40),
		Password: plainTextPassword,
	}

	// Call the normal create user flow.
	createdUser, err := Create(user)
	if err != nil {
		fmt.Errorf("error creating user: %v", err)
		return User{}, err
	}

	// This hurts lol
	createdUser.Password = plainTextPassword

	return createdUser, nil
}
