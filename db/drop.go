package db

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

// DropDatabaseWithName will drop the Postgres database with the given name.
// This function assumes the following:
//   1. That the postgres user exists
//   2. That the postgres database exists
//   3. That Postgres is configured to accept local connections with the 'password' method
//   4. That the postgres user's password is 'postgres'
func DropDatabaseWithName(name string) error {
	params := ConnectionParams{
		Host:     viper.GetString("ROOT_DB_HOST"),
		Name:     viper.GetString("ROOT_DB_NAME"),
		User:     viper.GetString("ROOT_DB_USER"),
		Password: viper.GetString("ROOT_DB_PASSWORD"),
	}

	logrus.WithFields(logrus.Fields{
		"host":     viper.GetString("ROOT_DB_HOST"),
		"name":     viper.GetString("ROOT_DB_NAME"),
		"user":     viper.GetString("ROOT_DB_USER"),
		"password": viper.GetString("ROOT_DB_PASSWORD"),
		"connstr":  params.String(),
	}).Debug("Creating root database connection.")
	db, err := NewConnectionFromParams(params)
	if err != nil {
		return errors.WithMessage(err, "failed to connect to root postgres database")
	}

	q := fmt.Sprintf("DROP DATABASE %s;", name)
	logrus.WithFields(logrus.Fields{
		"host":     viper.GetString("ROOT_DB_HOST"),
		"name":     viper.GetString("ROOT_DB_NAME"),
		"user":     viper.GetString("ROOT_DB_USER"),
		"password": viper.GetString("ROOT_DB_PASSWORD"),
		"connstr":  params.String(),
		"query":    q,
	}).Debug("Executing drop database query.")
	if _, err := db.Exec(q); err != nil {
		return errors.WithMessage(err, "failed to drop database")
	}

	logrus.WithFields(logrus.Fields{
		"host":     viper.GetString("ROOT_DB_HOST"),
		"name":     viper.GetString("ROOT_DB_NAME"),
		"user":     viper.GetString("ROOT_DB_USER"),
		"password": viper.GetString("ROOT_DB_PASSWORD"),
		"connstr":  params.String(),
	}).Debug("Closing root connection.")
	if err := db.Close(); err != nil {
		return errors.WithMessage(err, "there was an error closing the database connection")
	}

	logrus.WithFields(logrus.Fields{
		"database": name,
	}).Info("Database dropped.")

	return nil
}

// DropDatabaseFromConfig drops the database identified by the project's YAML config file.
// This function wraps DropDatabaseWithName. See the documentation for that function for an enumeration of some of the
// assumptions that are made when connecting to Postgres.
func DropDatabaseFromConfig() {
	logrus.WithFields(logrus.Fields{
		"database":    viper.GetString("DB_NAME"),
		"environment": viper.GetString("LAIR_ENV"),
	}).Debug("Will create database from environment config.")
	if err := DropDatabaseWithName(viper.GetString("DB_NAME")); err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			logrus.WithFields(logrus.Fields{
				"database":    viper.GetString("DB_NAME"),
				"environment": viper.GetString("LAIR_ENV"),
			}).Warn("Not dropping database because it does not exist.")
		} else {
			logrus.WithFields(logrus.Fields{
				"error": errors.WithStack(err),
			}).Panic("Encountered an unexpected error while dropping database.")
		}
	}
}
