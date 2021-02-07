package server

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type DBConfig struct {
	URL   string

	Logger *log.Logger
}

func (d *DBConfig) NewDB(ctx context.Context) (*sql.DB, error) {

	if d.URL == "" {
		return nil, errors.New("no db url")
	}



	db, err := sql.Open("postgres", d.URL)

	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("no db")
	}

	return db, nil

}

