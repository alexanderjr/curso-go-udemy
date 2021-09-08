package mysql

import (
	"api/src/domain/users/entity"
	userDomain "api/src/domain/users/service"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type UserMySQLRepository struct {
	db *sql.DB
}

const AlreadyExistsCode = 1062

func NewUserMySQLRepository(db *sql.DB) *UserMySQLRepository {
	return &UserMySQLRepository{db}
}

func (r UserMySQLRepository) Create(user entity.User) (*entity.User, error) {
	statement, err := r.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if !ok {
			return nil, err
		}

		if me.Number == AlreadyExistsCode {
			return nil, userDomain.ErrUserAlreadyExists
		}

		return nil, err
	}

	ultimoIdInserido, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = uint64(ultimoIdInserido)
	return &user, nil
}

func (r UserMySQLRepository) GetAll() ([]entity.User, error) {
	//TODO: add filter
	linhas, erro := r.db.Query(
		"SELECT id, nome, nick, email, criadoEm from usuarios",
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()
	var users []entity.User

	for linhas.Next() {
		var u entity.User
		if erro = linhas.Scan(
			&u.ID,
			&u.Name,
			&u.Nick,
			&u.Email,
			&u.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, u)
	}

	return users, nil
}
