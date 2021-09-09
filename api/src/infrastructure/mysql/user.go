package mysql

import (
	"api/src/domain/users/entity"
	"api/src/domain/users/service"
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
	rows, err := r.db.Query(
		"SELECT id, nome, nick, email, criadoEm from usuarios",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []entity.User

	for rows.Next() {
		var u entity.User
		if err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Nick,
			&u.Email,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r UserMySQLRepository) FindById(id int) (*entity.User, error) {
	rows, err := r.db.Query(
		"SELECT id, nome, nick, email, criadoEm from usuarios WHERE id = ?", id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	rows.Scan()
	var u entity.User
	hasUser := rows.Next()

	if !hasUser {
		return nil, userDomain.ErrUserNotFound
	}

	if err = rows.Scan(
		&u.ID,
		&u.Name,
		&u.Nick,
		&u.Email,
		&u.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r UserMySQLRepository) Delete(id int) error {
	statement, err := r.db.Prepare("DELETE from usuarios WHERE id = ?")

	if err != nil {
		return err
	}

	result, err := statement.Exec(id)

	res, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if res == 0 {
		return service.ErrUserNotFound
	}

	statement.Close()

	return nil
}
