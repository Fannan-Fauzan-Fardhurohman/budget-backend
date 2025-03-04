package services

import (
	"bougette-backend/cmd/api/requests"
	"bougette-backend/common"
	"bougette-backend/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (userService UserService) RegisterUser(userRegister *requests.RegisterUserRequest) (*models.UserModel, error) {
	hashedPassword, err := common.HashPassword(userRegister.Password)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("registration failed")
	}
	createdUser := models.UserModel{
		FirstName: &userRegister.FirstName,
		LastName:  &userRegister.LastName,
		Email:     userRegister.Email,
		Password:  hashedPassword,
	}
	result := userService.db.Create(&createdUser)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, errors.New("registration Failed")
	}
	fmt.Println(hashedPassword)
	return &createdUser, nil
}

func (userService UserService) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	result := userService.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
