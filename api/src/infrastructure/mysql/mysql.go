package mysql

import (
	"api/src/infrastructure"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() (*sql.DB, error) {
	db, erro := sql.Open("mysql", infrastructure.ConnectionStringDatabase)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
