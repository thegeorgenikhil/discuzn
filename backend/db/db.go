package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var DB *sql.DB

func ConnectDB(connectionString string) (*sql.DB, error) {
	result, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	err = result.Ping()

	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB!")
	return result, nil
}
