package database

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/steffen25/golang.zone/config"
)

type DB struct {
	*sql.DB
}

func NewDB(dbCfg config.MysqlConfig) (*DB, error) {
	//DSN := fmt.Sprintf("%s:%s@unix(/tmp/mysql.sock)/%s?parseTime=true", dbCfg.Username, dbCfg.Password, dbCfg.DatabaseName)
	dataSourceName := fmt.Sprintf("%s:%s@/%s?parseTime=true", dbCfg.Username, dbCfg.Password, dbCfg.DatabaseName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}