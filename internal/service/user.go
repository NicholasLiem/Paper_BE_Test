package service

import (
	"github.com/NicholasLiem/Paper_BE_Test/internal/datastruct"
	"github.com/NicholasLiem/Paper_BE_Test/internal/repository"
	"github.com/NicholasLiem/Paper_BE_Test/utils"
	"log"
	"net/http"
)

type UserService interface {
	CreateUser(user datastruct.User) (bool, *utils.HttpError)
	GetUser(userID uint) (*datastruct.User, *utils.HttpError)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (us *userService) CreateUser(user datastruct.User) (bool, *utils.HttpError) {
	if user.Name == "" || user.Email == "" {
		return false, &utils.HttpError{Message: "Invalid user data", StatusCode: http.StatusBadRequest}
	}

	success, err := us.dao.NewUserQuery().CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return false, &utils.HttpError{Message: "Error creating user", StatusCode: http.StatusInternalServerError}
	}
	return success, nil
}

func (us *userService) GetUser(userID uint) (*datastruct.User, *utils.HttpError) {
	user, err := us.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return nil, &utils.HttpError{Message: "Error retrieving user", StatusCode: http.StatusInternalServerError}
	}
	return user, nil
}
