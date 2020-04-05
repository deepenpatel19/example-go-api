package postgres

import (
	"fmt"
	"github.com/cobbinma/example-go-api/pkg/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DBClient interface {
}

type dbClient struct {
	db *sqlx.DB
}

func NewDBClient() DBClient {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		config.DBHost,
		config.DBName,
		config.DBUser,
		config.DBPassword)

	driver := "postgres"

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logrus.Fatalln("Could not open database: ", err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return &dbClient{db: db}
}
