package service

import (
	"errors"
	"quiz/dto"
	"quiz/models"
	"quiz/repository"

	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	RegisterUser(input dto.RegisterUserDto) (*models.User, error)
	Login(input dto.LoginDto) (*models.User, error)
	IsEmaillAvailabilty(input string) (bool, error)
	GetUserByid(ID int) (*models.User, error)
	DeleteUser(id int) (*models.User, error)
	// SaveAvatar(ID int, fileLocation string) (User, error)
}

type service_user struct {
	repository_user repository.RepositoryUser
}

func NewServiceUser(repository_user repository.RepositoryUser) *service_user {
	return &service_user{repository_user}
}

func (s *service_user) RegisterUser(input dto.RegisterUserDto) (*models.User, error) {
	user := &models.User{}

	user.Name = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.Role = 0
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository_user.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service_user) Login(input dto.LoginDto) (*models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository_user.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("user not found that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service_user) DeleteUser(id int) (*models.User, error) {
	user, err := s.repository_user.FindById(id)
	if err != nil {
		return user, err
	}
	userDel, err := s.repository_user.Delete(user)

	if err != nil {
		return userDel, err
	}
	return userDel, nil
}

func (s *service_user) IsEmaillAvailabilty(input string) (bool, error) {
	//karena hanya email maka di inisiasi hanya email
	emailUser := &models.User{}
	emailUser.Email = input

	//pengambilan algoritmanya repository yaitu findbyemail
	user, err := s.repository_user.FindByEmail(input)
	if err != nil {
		return false, err
	}

	// ini nilainya true karena misal kita input email ini sama ga dengan email yang terdaftar dg id sekian
	//kalau g ada maka balikkanya 0 sehingga bisa di daftrakan atau availabilty
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service_user) GetUserByid(ID int) (*models.User, error) {
	user, err := s.repository_user.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user Not Found With That ID")
	}

	return user, nil

}
