package database

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // load the Postgres driver for golang-migrate/migrate
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func newMigration() (*migrate.Migrate, error) {
	db, err := NewConnectionFromConfig()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to connect to database")
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, errors.WithMessage(err, "could not get an instance driver")
	}
	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", viper.GetString("MIGRATIONS_DIRECTORY")),
		"postgres",
		driver,
	)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create a new migration")
	}

	return migration, nil
}

func CreateNewMigrationFile(desc string) {
	// Get a sorted list of templates in the project migrations directory
	f, err := ioutil.ReadDir("./database/migrations")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to read project migrations directory.")
	}

	// Get the first six characters of the last filename in the sorted list
	// The first six characters of the file should be the migration version
	v := []rune(f[len(f)-1].Name())[:6]

	// Convert the migration version to an integer
	current, err := strconv.Atoi(string(v))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"latest_migration_file": f[len(f)-1],
		}).Warn("No migration templates found. Creating version 000001.")
		current = 0
	} else {
		logrus.WithFields(logrus.Fields{
			"current_version":       current,
			"next_version":          current + 1,
			"latest_migration_file": f[len(f)-1],
		}).Info("Latest migration file found.")
	}

	// Reformat the description so it is lower case and spaces are replaced with underscores
	d := strings.ToLower(strings.Replace(desc, " ", "_", -1))

	// Build the filenames for the up and down migrations
	down := fmt.Sprintf("database/migrations/%06d_%s.down.sql", current+1, d)
	up := fmt.Sprintf("database/migrations/%06d_%s.up.sql", current+1, d)

	// Create the up and down migration templates
	logrus.WithFields(logrus.Fields{
		"filename": down,
	}).Info("Creating down migration.")
	if err := ioutil.WriteFile(down, []byte{}, 0664); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": errors.WithStack(err),
		}).Fatal("Failed to create down migration.")
	}

	logrus.WithFields(logrus.Fields{
		"filename": up,
	}).Info("Creating up migration.")
	if err := ioutil.WriteFile(up, []byte{}, 0664); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": errors.WithStack(err),
		}).Fatal("Failed to create up migration.")
	}
}

func GetDatabaseVersion(env string) {
	migration, err := newMigration()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"environment": env,
			"error":       errors.WithStack(err),
		}).Fatal("Failed to create a new migration.")
	}

	v, d, err := migration.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			fmt.Println("No migrations have been applied")
		}
	} else {
		fmt.Printf(
			"Current version: %d\nDirty: %t\n",
			v,
			d,
		)
	}
}

func MigrateDatabase(steps int) {
	migration, err := newMigration()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"database": viper.GetString("DB_NAME"),
			"steps":       steps,
			"error":       errors.WithStack(err),
		}).Fatal("Failed to create a new migration.")
	}

	// FIXME: "no change" error should not cause a panic
	// FIXME: No migration files should not cause a fatal error
	if steps == 0 {
		logrus.WithFields(logrus.Fields{
			"database": viper.GetString("DB_NAME"),
		}).Info("Migrating database to latest version.")
		if err := migration.Up(); err != nil {
			logrus.WithFields(logrus.Fields{
				"database": viper.GetString("DB_NAME"),
				"steps":       steps,
				"error":       errors.WithStack(err),
			}).Fatal("Database migration failed.")
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"database": viper.GetString("DB_NAME"),
			"steps":       steps,
		}).Info("Performing step-wise migration on database.")
		if err := migration.Steps(steps); err != nil {
			logrus.WithFields(logrus.Fields{
				"database": viper.GetString("DB_NAME"),
				"steps":       steps,
				"error":       errors.WithStack(err),
			}).Fatal("Database migration failed.")
		}
	}
}
