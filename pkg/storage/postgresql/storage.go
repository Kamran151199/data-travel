package postgresql

import (
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Storage is a struct with minimalistic fields to represent a SQL storage.
// It is used to create a SQL storage client.
type Storage struct {
	Host     string `json:"host" validate:"required" default:"localhost"`
	Database string `json:"database" validate:"required" default:"core"`
	User     string `json:"user" validate:"required" default:"postgres"`
	Password string `json:"password" validate:"required" default:"password"`
	Port     string `json:"port" validate:"required" default:"5432"`
	MaxConn  int    `json:"max_conn" default:"200"`
}

// GetDB does what its name implies
func (storage *Storage) GetDB() (*sqlx.DB, error) {
	// First set up the pgx connection pool
	connConfig := pgx.ConnConfig{
		Host:     storage.Host,
		Database: storage.Database,
		User:     storage.User,
		Password: storage.Password,
	}

	// Set up the connection pool
	connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		MaxConnections: storage.MaxConn,
	})

	// If there was an error creating the connection pool, return an error
	if err != nil {
		return nil, errors.Wrap(err, "Call to pgx.NewConnPool failed")
	}

	// Then set up sqlx and return the created DB reference
	nativeDB := stdlib.OpenDBFromPool(connPool)

	// if there was an error creating the native DB, return an error
	if err != nil {
		connPool.Close()
		return nil, errors.Wrap(err, "Call to stdlib.OpenFromConnPool failed")
	}

	return sqlx.NewDb(nativeDB, "pgx"), nil
}
