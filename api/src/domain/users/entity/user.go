package entity

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
)

var (
	ErrInvalidName  = errors.New("o nome é obrigatório")
	ErrInvalidNick  = errors.New("o nick é obrigatório")
	ErrInvalidEmail = errors.New("o email é obrigatório")
)

type User struct {
	ID        uint64
	Name      string
	Nick      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u User) IsValid() error {
	if u.Name == "" {
		return ErrInvalidName
	}

	if u.Nick == "" {
		return ErrInvalidNick
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return ErrInvalidEmail
	}

	if u.Password == "" {
		return errors.New("a senha é obrigatório")
	}

	return nil
}
