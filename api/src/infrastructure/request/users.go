package request

import "time"

type UserRequest struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty`
	Nick     string    `json:"nick,omitempty`
	Email    string    `json:"email,omitempty`
	Senha    string    `json:"senha,omitempty`
	CriadoEm time.Time `json:"CriadoEm,omitempty`
}
