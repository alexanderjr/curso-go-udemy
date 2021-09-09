package repository

import "api/src/domain/users/entity"

type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
	GetAll() ([]entity.User, error)
	FindById(id int) (*entity.User, error)
	Delete(id int) error
}
