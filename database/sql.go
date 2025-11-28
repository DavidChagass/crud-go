package database

import (
	"database/sql"
	"fmt"
	//importa o mysql
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {

	dbconn := "root:12345@tcp(127.0.0.1:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dbconn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a conexão ao banco: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("erro ao pingar com o banco de dados:  %w", err)
	}
	return db, nil
}
