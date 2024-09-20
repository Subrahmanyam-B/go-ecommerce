package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"log"
)

type UserService struct {
	// UserService struct
}

func (s UserService) SignUp(input dto.UserSignup) (string, error) {

	log.Println(input)

	return "example token", nil
}

func (s UserService) findUserbyEmail(email string) (*domain.User, error) {
	// perform some DB operation
	// and some business logic
	return nil, nil
}

func (s UserService) Login(input any) (string, error) {
	return "", nil
}

func (s UserService) GetVerificationCode(user domain.User) (int, error) {
	return 0, nil
}

func (s UserService) verifyCode(id uint, code int) error {

	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint, input any) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {

	return nil, nil
}

func (s UserService) GetOrderbyId(id uint, userId uint) (interface{}, error) {

	return nil, nil
}
