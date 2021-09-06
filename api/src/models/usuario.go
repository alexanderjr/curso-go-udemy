package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty`
	Nick     string    `json:"nick,omitempty`
	Email    string    `json:"email,omitempty`
	Senha    string    `json:"senha,omitempty`
	CriadoEm time.Time `json:"CriadoEm,omitempty`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatório")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return erro
	}

	if etapa == "atualizar" {
		return nil
	}

	if usuario.Senha == "" {
		return errors.New("a senha é obrigatório")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Nick = strings.TrimSpace(usuario.Nick)

	if etapa == "cadastrar" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)

		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}
	return nil
}

func (usuario *Usuario) Atualizar(usuarioAtualizado Usuario) {
	usuario.Nome = usuarioAtualizado.Nome
	usuario.Email = usuarioAtualizado.Email
	usuario.Nick = usuarioAtualizado.Nick
}
