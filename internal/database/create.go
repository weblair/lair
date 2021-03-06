package database

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

// TODO: Add debug logging to create.go

// CreateDatabaseWithName will create a new Postgres database with the given name.
// This function will drop the database using the creds in the config under ROOT_DB_HOST, ROOT_DB_NAME, ROOT_DB_USER,
// and ROOT_DB_PASSWORD.
func CreateDatabaseWithName(name string, force bool) error {
	params := ConnectionParams{
		Host:     viper.GetString("ROOT_DB_HOST"),
		Name:     viper.GetString("ROOT_DB_NAME"),
		User:     viper.GetString("ROOT_DB_USER"),
		Password: viper.GetString("ROOT_DB_PASSWORD"),
	}

	db, err := NewConnectionFromParams(params)
	if err != nil {
		return errors.WithMessage(err, "failed to connect to root postgres database")
	}

	if force {
		if err := DropDatabaseWithName(name); err != nil {
			if strings.Contains(err.Error(), "does not exist") {
				logrus.WithFields(logrus.Fields{
					"database": viper.GetString("DB_NAME"),
				}).Warn("Not dropping database because it does not exist.")
			} else {
				return errors.WithMessagef(err, "could not drop existing database %s", name)
			}
		}
	}
	if _, err := db.Exec(fmt.Sprintf("CREATE DATABASE %s;", name)); err != nil {
		return errors.WithMessagef(err, "failed to create database %s", name)
	}

	if err := db.Close(); err != nil {
		return errors.WithMessage(err, "there was an error closing the database connection")
	}

	logrus.WithFields(logrus.Fields{
		"database": name,
	}).Info("Database created.")

	return nil
}

// CreateDatabaseFromConfig creates a database using the name from the project's YAML config file.
// This function wraps CreateDatabaseWithName. See the documentation for that function for an enumeration of some of the
// assumptions that are made when connecting to Postgres.
func CreateDatabaseFromConfig(force bool) {
	if err := CreateDatabaseWithName(viper.GetString("DB_NAME"), force); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			logrus.WithFields(logrus.Fields{
				"database":    viper.GetString("DB_NAME"),
			}).Warn("Skipping creation of database that already exists.")
		} else {
			logrus.WithFields(logrus.Fields{
				"error": errors.WithStack(err),
			}).Panic("Encountered an unexpected error while creating database.")
		}
	}
}
