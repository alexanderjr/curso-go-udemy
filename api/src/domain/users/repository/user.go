package repository

import (
	users "api/src/domain/users/entity"
)

type UserRepository interface {
	Create(user users.User) (*users.User, error)
}
