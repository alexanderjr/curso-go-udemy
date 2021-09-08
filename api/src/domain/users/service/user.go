package service

import (
	"api/src/domain/users/entity"
	"api/src/domain/users/repository"
	"errors"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type UserService struct {
	repository repository.UserRepository
}

func (u UserService) Create(user entity.User) (*entity.User, error) {
	if err := user.IsValid(); err != nil {
		return nil, err
	}

	userAdded, err := u.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return userAdded, nil
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{repository: r}
}
