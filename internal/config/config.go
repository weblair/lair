package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func InitLogging() error {
	// Valid log levels are:
	//	- panic
	//	- fatal
	//	- error
	//	- warn
	//	- info
	//	- debug
	//	- trace
	loglevel, err := logrus.ParseLevel(viper.GetString("LOGLEVEL"))
	// TODO: Review error handling for initLogging
	if err != nil {
		return errors.WithStack(err)
	}

	logrus.SetLevel(loglevel)

	return nil
}

// LoadEnvConfig will load the given environment from its corresponding YAML config file.
// Use this to override the default environment that is loaded from the .env file.
func LoadEnvConfig(env string) {
	var dbHost, dbUser, dbName, dbPassword string
	if viper.GetString("DB_HOST") != "localhost" {
		dbHost = viper.GetString("DB_HOST")
	}
	if viper.GetString("DB_USER") != "postgres" {
		dbUser = viper.GetString("DB_USER")
	}
	if viper.GetString("DB_NAME") != "" {
		dbName = viper.GetString("DB_NAME")
	}
	if viper.GetString("DB_PASSWORD") != "postgres" {
		dbPassword = viper.GetString("DB_PASSWORD")
	}

	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.SetConfigName(env)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logrus.WithFields(logrus.Fields{
				"environment": env,
				"error":       errors.WithStack(err),
			}).Fatal("Failed to read configuration file.")
		}
	}

	if dbHost != "" {
		viper.Set("DB_HOST", dbHost)
	}
	if dbUser != "" {
		viper.Set("DB_USER", dbUser)
	}
	if dbName != "" {
		viper.Set("DB_NAME", dbName)
	}
	if dbPassword != "" {
		viper.Set("DB_PASSWORD", dbPassword)
	}
}

// TODO: Set up configuration for loglevel
// TODO: Refactor init to avoid superfluous initialization of config values.
func init() {
	// Default config values
	viper.SetDefault("ROOT_DB_HOST", "localhost")
	viper.SetDefault("ROOT_DB_USER", "postgres")
	viper.SetDefault("ROOT_DB_NAME", "postgres")
	viper.SetDefault("ROOT_DB_PASSWORD", "postgres")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("MIGRATIONS_DIRECTORY", "migrations")
	viper.SetDefault("SEED_DIRECTORY", "seed")

	_ = gotenv.Load()
	viper.AutomaticEnv()

}
