package repository

import (
	"errors"
	"github.com/NicholasLiem/Paper_BE_Test/internal/datastruct"
	"gorm.io/gorm"
)

type UserQuery interface {
	BeginTransaction() *gorm.DB
	CreateUser(user datastruct.User) (bool, error)
	CreateUserTx(user datastruct.User, tx *gorm.DB) (bool, error)
	UpdateUserTx(userID uint, updatedUser datastruct.User, tx *gorm.DB) (bool, error)
	DeleteUser(userID uint) (bool, error)
	GetUser(userID uint) (*datastruct.User, error)
	GetAllUsers() ([]datastruct.User, error)
	FindUserByEmail(email string) (*datastruct.User, error)
}

type userQuery struct {
	pgdb *gorm.DB
}

func NewUserQuery(db *gorm.DB) UserQuery {
	return &userQuery{pgdb: db}
}

func (uq *userQuery) BeginTransaction() *gorm.DB {
	return uq.pgdb.Begin()
}

func (uq *userQuery) CreateUser(user datastruct.User) (bool, error) {
	result := uq.pgdb.Create(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (uq *userQuery) CreateUserTx(user datastruct.User, tx *gorm.DB) (bool, error) {
	result := tx.Create(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (uq *userQuery) DeleteUser(userID uint) (bool, error) {
	result := uq.pgdb.Delete(&datastruct.User{}, userID)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (uq *userQuery) UpdateUserTx(userID uint, updatedUser datastruct.User, tx *gorm.DB) (bool, error) {
	var user datastruct.User
	if err := tx.First(&user, userID).Error; err != nil {
		return false, err
	}

	result := tx.Model(&user).Updates(updatedUser)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (uq *userQuery) GetUser(userID uint) (*datastruct.User, error) {
	var user datastruct.User
	result := uq.pgdb.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (uq *userQuery) GetAllUsers() ([]datastruct.User, error) {
	var users []datastruct.User
	result := uq.pgdb.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (uq *userQuery) FindUserByEmail(email string) (*datastruct.User, error) {
	var user datastruct.User
	if err := uq.pgdb.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
