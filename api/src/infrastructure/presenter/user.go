package presenter

import (
	"api/src/domain/users/entity"
	"strconv"
)

func ShowUser(u entity.User) interface{} {
	return struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		Nome      string `json:"nome"`
		Nick      string `json:"nick"`
		CreatedAt string `json:"created_at"`
	}{
		ID:        strconv.Itoa(int(u.ID)),
		Nome:      u.Name,
		Email:     u.Email,
		Nick:      u.Nick,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
