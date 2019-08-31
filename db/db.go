package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // load the Postgres driver for database/sql
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ConnectionParams is the collection of parameters needed to build a connection string.
type ConnectionParams struct {
	Host     string
	Name     string
	User     string
	Password string
}

// String builds a Postgres connection string from the ConnectionParams struct.
func (c ConnectionParams) String() string {
	return fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		c.Host,
		c.Name,
		c.User,
		c.Password,
	)
}

// NewConnectionFromParams creates a new database connection from the given parameters.
func NewConnectionFromParams(c ConnectionParams) (*sql.DB, error) {
	logrus.WithFields(logrus.Fields{
		"host":              c.Host,
		"name":              c.Name,
		"user":              c.User,
		"password":          c.Password,
		"connection_string": c.String(),
	}).Debug("Creating new database connection.")
	db, err := sql.Open("postgres", c.String())
	if err != nil {
		return nil, errors.WithMessage(err, "failed to connect to database")
	}

	return db, nil
}

// NewConnectionFromConfig creates a new database connection based on the project's configuration templates.
func NewConnectionFromConfig() (*sql.DB, error) {
	params := ConnectionParams{
		Host:     viper.GetString("DB_HOST"),
		Name:     viper.GetString("DB_NAME"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
	}

	return NewConnectionFromParams(params)
}
