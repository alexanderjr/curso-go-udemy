package request

import (
	"api/src/domain/users/entity"
	"time"
)

type UserRequest struct {
	Nome  string `json:"nome,omitempty`
	Nick  string `json:"nick,omitempty`
	Email string `json:"email,omitempty`
	Senha string `json:"senha,omitempty`
}

func (u UserRequest) CreateDomainUser() entity.User {
	return entity.User{
		Name:      u.Nome,
		Email:     u.Email,
		Nick:      u.Nick,
		Password:  u.Senha,
		CreatedAt: time.Now(),
	}
}

func (u UserRequest) CreateDomainUserToUpdate(id uint64) entity.User {
	return entity.User{
		ID:        id,
		Name:      u.Nome,
		Email:     u.Email,
		Nick:      u.Nick,
		Password:  u.Senha,
		CreatedAt: time.Now(),
	}
}
