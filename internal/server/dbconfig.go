package server

import (
"context"
"database/sql"
"errors"
"fmt"
_ "github.com/jinzhu/gorm/dialects/postgres"
log "github.com/sirupsen/logrus"
"time"
)

type DBConfig struct {
	Host               string
	Port               int
	User               string
	Pass               string
	DBName             string
	SSLMode            string
	SSLCert            string
	SSLKey             string
	SSLServerCA        string
	LoggingEnabled     bool
	Dialect            string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    time.Duration

	Logger *log.Logger
}

func (d *DBConfig) NewDB(ctx context.Context) (*sql.DB, error) {

	if d.User == "" {
		return nil, errors.New("no db user")
	}
	if d.Pass == "" {
		return nil, errors.New("no db pass")
	}
	if d.Port == 0 {
		return nil, errors.New("no db port")
	}
	if d.DBName == "" {
		return nil, errors.New("no db name")
	}
	if d.SSLMode == "" {
		return nil, errors.New("no db ssl mode")
	}
	if d.SSLCert == "" {
		return nil, errors.New("no db cert")
	}
	if d.SSLKey == "" {
		return nil, errors.New("no db key")
	}
	if d.SSLServerCA == "" {
		return nil, errors.New("no db ssl ca")
	}

	if d.Dialect == "" {
		return nil, errors.New("no db dialect")
	}

	source := d.DSN()
	db, err := sql.Open("postgres", source)

	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("no db")
	}
	// Set the maximum number of concurrently open connections (in-use + idle)
	// to 5. Setting this to less than or equal to 0 will mean there is no
	// maximum limit (which is also the default setting).
	db.SetMaxOpenConns(d.MaxOpenConnections)

	// Set the maximum number of concurrently idle connections to 5. Setting this
	// to less than or equal to 0 will mean that no idle connections are retained.
	db.SetMaxIdleConns(d.MaxIdleConnections)

	// Set the maximum lifetime of a connection to 1 hour. Setting it to 0
	// means that there is no maximum lifetime and the connection is reused
	// forever (which is the default behavior).
	db.SetConnMaxLifetime(d.ConnMaxLifetime)

	return db, nil

}
func (d DBConfig) DSN() string {

	cs := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s ", d.Host, d.Port, d.User, d.Pass, d.DBName, d.SSLMode, d.SSLCert, d.SSLKey, d.SSLServerCA)

	d.Logger.Debug(cs)
	return cs
}

