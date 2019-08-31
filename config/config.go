package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
	"os/user"
)

// TODO: Don't create DB_*, GIN_MODE, or LAIR_ENV variables in a new root config.
// TODO: Use uppercase keys when creating the root config file.
func createRootConfig(homeDir string) error {
	logrus.Info("Creating .lair directory")
	if err := os.MkdirAll(homeDir+"/.lair/", 0755); err != nil {
		return errors.WithMessage(err, "failed to create .lair directory")
	}

	logrus.Info("Creating root config")
	if err := viper.WriteConfigAs(homeDir + "/.lair/lairrc.yml"); err != nil {
		return errors.WithMessage(err, "failed to create root configuration file")
	}

	return nil
}

func initRootConfig() error {
	u, err := user.Current()
	if err != nil {
		return errors.WithMessage(err, "failed to get current user")
	}

	viper.SetDefault("ROOT_LOGLEVEL", "info")
	viper.SetDefault("ROOT_DB_HOST", "localhost")
	viper.SetDefault("ROOT_DB_USER", "postgres")
	viper.SetDefault("ROOT_DB_NAME", "postgres")
	viper.SetDefault("ROOT_DB_PASSWORD", "postgres")
	viper.SetDefault("GITHUB_USER", "weblair")
	viper.SetDefault("COPYRIGHT", "Robert Hawk")

	viper.AddConfigPath(u.HomeDir + "/.lair/")
	viper.SetConfigName("lairrc")
	if _, err := os.Stat(u.HomeDir + "/.lair/lairrc.yml"); err == nil {
		if err := viper.ReadInConfig(); err != nil {
			return errors.WithMessage(err, "failed to read root configuration file")
		}
	} else if os.IsNotExist(err) {
		logrus.Warn("Root config file not found. Generating a default config.")
		if err := createRootConfig(u.HomeDir); err != nil {
			return errors.WithStack(err)
		}
	} else if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func initLogging() error {

	// Valid log levels are:
	//	- panic
	//	- fatal
	//	- error
	//	- warn
	//	- info
	//	- debug
	//	- trace
	loglevel, err := logrus.ParseLevel(viper.GetString("ROOT_LOGLEVEL"))
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
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("LAIR_ENV", "development")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "")
	viper.SetDefault("DB_PASSWORD", "postgres")

	_ = gotenv.Load()
	viper.AutomaticEnv()

	viper.AddConfigPath("./config")
	viper.SetConfigName(viper.GetString("LAIR_ENV"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logrus.WithFields(logrus.Fields{
				"environment": viper.GetString("LAIR_ENV"),
				"error":       errors.WithStack(err),
			}).Fatal("Failed to read configuration file.")
		}
	}
}

// TODO: Set up configuration for loglevel
// TODO: Refactor init to avoid superfluous initialization of config values.
func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)

	if err := initRootConfig(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to load root configuration.")
	}

	if err := initLogging(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to initialize logger.")
	}
}
