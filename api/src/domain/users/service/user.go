package service

import (
	"api/src/domain/users/contract"
	"api/src/domain/users/entity"
	"api/src/domain/users/repository"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found exists")
)

type UserService struct {
	repository repository.UserRepository
	encriptor  contract.Encryptor
}

func NewUserService(r repository.UserRepository, e contract.Encryptor) UserService {
	return UserService{repository: r, encriptor: e}
}

func (u UserService) Create(user entity.User) (*entity.User, error) {
	if err := user.IsValid(); err != nil {
		return nil, err
	}

	hashValue, err := u.encriptor.ToHash(user.Password)
	if err != nil {
		return nil, err
	}

	user.ToHashPassword(string(hashValue))

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

func (u UserService) Update(user entity.User) (*entity.User, error) {
	if err := user.IsValid(); err != nil {
		return nil, err
	}

	if err := u.repository.Update(user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserService) FindById(id int) (*entity.User, error) {
	user, err := u.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) Delete(id int) error {
	if err := u.repository.Delete(id); err != nil {
		return err
	}

	return nil
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
