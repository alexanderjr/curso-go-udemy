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

func (u UserService) GetAll() ([]entity.User, error) {
	users, err := u.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func NewUserService(r repository.UserRepository) UserService {
	return UserService{repository: r}
}

func (u UserService) ErrFromDomain(err error) bool {
	domainErrors := []error{
		ErrUserAlreadyExists,
		entity.ErrInvalidEmail,
		entity.ErrInvalidName,
		entity.ErrInvalidNick,
	}

	for i := 0; i < len(domainErrors); i++ {
		if errors.Is(err, domainErrors[i]) {
			return true
		}
	}

	return false
}
