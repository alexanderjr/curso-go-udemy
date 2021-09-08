package presenter

import (
	"api/src/domain/users/entity"
	"strconv"
)

type UserPresenter struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Nome      string `json:"nome"`
	Nick      string `json:"nick"`
	CreatedAt string `json:"created_at"`
}

func ShowUser(u entity.User) interface{} {
	return UserPresenter{
		ID:        strconv.Itoa(int(u.ID)),
		Nome:      u.Name,
		Email:     u.Email,
		Nick:      u.Nick,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ShowAllUser(u []entity.User) interface{} {
	var users []interface{}

	for i := 0; i < len(u); i++ {
		users = append(users, ShowUser(u[i]))
	}

	return users
}
