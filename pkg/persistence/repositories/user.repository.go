package repositories

import (
	"context"
	"fmt"
	"log"

	"training/tutorialGL/pkg/domain/schemas"
	"training/tutorialGL/pkg/interfaces"
)

type UserRepository struct {
	Database interfaces.DatabaseInterface
}

const (
	colUsername = "username"
	colFullname = "fullname"
	colEmail    = "email"

	tableUsers = "users"
)

var userCols = []string{colUsername, colFullname, colEmail}

func (u UserRepository) Insert(ctx context.Context, user schemas.UserSchema) error {

	newUser := schemas.NewUserSchema(user)

	log.Println("new User", newUser.StructToMap())

	err := u.Database.Insert(ctx, newUser.StructToMap(), tableUsers)

	if err != nil {
		log.Fatal(":::::::::::::::::ERROR REPO:::::::::::::::::", err)
		return err
	}

	return nil
}

func (u UserRepository) Get(ctx context.Context) ([]schemas.UserSchema, error) {

	var user schemas.UserSchema

	rows, err := u.Database.Select(ctx, userCols, tableUsers)

	if err != nil {
		// if any error occurred while reading rows.
		fmt.Println("Error will reading user table: ", err)
		return nil, err
	}

	defer rows.Close()
	var userSlice []schemas.UserSchema

	for rows.Next() {
		rows.Scan(&user.Username, &user.Fullname, &user.Email)
		userSlice = append(userSlice, user)
	}

	return userSlice, err
}
