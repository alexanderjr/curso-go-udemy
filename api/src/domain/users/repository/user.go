package repository

import "api/src/domain/users/entity"

type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
}
