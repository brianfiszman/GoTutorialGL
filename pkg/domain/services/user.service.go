package services

import (
	"context"
	"encoding/json"
	"log"
	"training/tutorialGL/pkg/application/dtos"
	"training/tutorialGL/pkg/domain/schemas"
	"training/tutorialGL/pkg/persistence/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (u UserService) Create(ctx context.Context, createUserDTO dtos.CreateUserDTO) (dtos.CreateUserDTOResponse, error) {

	user := userDTOToUserSchema(createUserDTO)

	err := u.Repository.Insert(ctx, user)

	if err != nil {
		log.Fatal(":::::::::::::::::::::::::::ERROR SERVICE::::::::::::::::::::::::", err)
	}

	var responseUserDTO = dtos.CreateUserDTOResponse{
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email,
	}

	return responseUserDTO, err
}

func (u UserService) Get(ctx context.Context) ([]dtos.GetUsersDTOResponse, error) {

	userSlice, err := u.Repository.Get(ctx)

	if err != nil {
		log.Fatal(":::::::::::::::::::::::::::ERROR SERVICE::::::::::::::::::::::::", err)
	}

	j, _ := json.Marshal(userSlice)

	var users []dtos.GetUsersDTOResponse

	json.Unmarshal([]byte(j), &users)

	return users, err
}

func userDTOToUserSchema(u dtos.CreateUserDTO) schemas.UserSchema {
	user := schemas.UserSchema{
		Username: u.Username,
		Fullname: u.Fullname,
		Email:    u.Email,
	}
	return user
}
