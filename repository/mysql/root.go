package mysql

import (
	"database/sql"
	"eCommerce/config"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(cfg *config.Config) (*MySQL, error) {
	m := new(MySQL)
	var err error

	if m.DB, err = sql.Open("mysql", cfg.MySQL.Uri); err != nil {
		return nil, err
	} else if err = m.DB.Ping(); err != nil {
		return nil, err
	} else {
		return m, nil
	}
}
