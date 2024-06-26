package infra

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gaspecian/go-to-do-list/config"
)

func DbConnect() *sql.DB {
	dbConfig := config.LoadEnv().DbConnection

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error while connecting to DB: %s", err)
	}

	return db
}
