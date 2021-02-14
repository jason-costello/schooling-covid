// Package server contains the needful to connect to the DB as well as startup the web server and provide the handlers.
package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ErrNoDBURL gets returned when a blank URL is provided the NewDB
var ErrNoDBURL = errors.New("no connection url provided for db")

// ErrFailedToOpenDB is returned on sql.Open error.  The original error gets
// wrapped in the return
var ErrFailedToOpenDB =	 errors.New("failed to open db")

// ErrNoDB is returned when sql.Open succeeds but the returned db is nil
var ErrNoDB = errors.New("no db found")

// DBConfig contains the url for the db.
type DBConfig struct {
	URL string
}

// NewDB verifies all of the needed configuration optinos are set
// and then attempts to open a connection to the database.  If
// unable to make a connection and have a valid db returned
// errors will be returned.
func (d *DBConfig) NewDB(ctx context.Context) (*sql.DB, error) {

	if d.URL == "" {
		return nil, ErrNoDBURL
	}

	db, err := sql.Open("postgres", d.URL)

	if err != nil {
		return nil, fmt.Errorf("server.NewDB: %s - %w", ErrFailedToOpenDB, err)
	}

	if db == nil {
		return nil, ErrNoDB
	}

	return db, nil

}
