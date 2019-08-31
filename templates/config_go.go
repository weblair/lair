package templates

// TODO: Add logging config

// ConfigGo is the template for the config/config.go file in a Gin project.
const ConfigGo = `package config

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func init() {
	_ = gotenv.Load()
	viper.AutomaticEnv()

	viper.SetDefault("LAIR_ENV", "development")
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("URL", "localhost:3000")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "$2_development")
	viper.SetDefault("DB_PASSWORD", "postgres")
}
`
