package repositorios

import (
	modelos "api/src/models"
	"database/sql"
	"errors"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm from usuarios WHERE nome LIKE ? OR nick LIKE ?",
		nomeOuNick,
		nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()
	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorID(id uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm from usuarios WHERE id = ?", id,
	)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()
	linhas.Scan()
	var usuario modelos.Usuario
	hasUser := linhas.Next()

	if !hasUser {
		return usuario, errors.New("User not found")
	}

	if erro = linhas.Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Nick,
		&usuario.Email,
		&usuario.CriadoEm,
	); erro != nil {
		return modelos.Usuario{}, erro
	}

	return usuario, nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email,senha, criadoEm from usuarios WHERE email = ?", email,
	)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()
	linhas.Scan()
	var usuario modelos.Usuario
	hasUser := linhas.Next()

	if !hasUser {
		return usuario, errors.New("User not found")
	}

	if erro = linhas.Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Nick,
		&usuario.Email,
		&usuario.Senha,
		&usuario.CriadoEm,
	); erro != nil {
		return modelos.Usuario{}, erro
	}

	return usuario, nil
}

func (repositorio Usuarios) DeletarPorID(id uint64) error {
	statement, erro := repositorio.db.Query(
		"DELETE from usuarios WHERE id = ?", id,
	)

	if erro != nil {
		return erro
	}

	statement.Close()

	return nil
}

func (repositorio Usuarios) Atualizar(usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Query(
		"UPDATE usuarios set nome = ?, email = ?, nick = ? WHERE id = ?",
		usuario.ID,
		usuario.Nome,
		usuario.Email,
		usuario.Nick,
	)

	if erro != nil {
		return erro
	}

	statement.Close()

	return nil
}
