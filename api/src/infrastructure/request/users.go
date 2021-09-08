package request

import (
	users "api/src/domain/users/entity"
	"time"
)

type UserRequest struct {
	Nome  string `json:"nome,omitempty`
	Nick  string `json:"nick,omitempty`
	Email string `json:"email,omitempty`
	Senha string `json:"senha,omitempty`
}

func (u UserRequest) CreateDomainUser() users.User {
	return users.User{
		Name:      u.Nome,
		Email:     u.Email,
		Nick:      u.Nick,
		Password:  u.Senha,
		CreatedAt: time.Now(),
	}
}
