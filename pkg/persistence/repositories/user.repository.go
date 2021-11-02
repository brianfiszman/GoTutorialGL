package repositories

import (
	"context"
	"log"
	"training/tutorialGL/pkg/domain/schemas"
	"training/tutorialGL/pkg/interfaces"
)

type UserRepository struct {
	Database interfaces.DatabaseInterface
}

const entityName string = "users"

func (u UserRepository) Insert(ctx context.Context, user schemas.UserSchema) error {

	newUser := schemas.NewUserSchema(user)

	err := u.Database.Insert(ctx, newUser.StructToMap(), entityName)

	if err != nil {
		log.Fatal(":::::::::::::::::ERROR REPO:::::::::::::::::", err)
		return err
	}

	return nil
}
