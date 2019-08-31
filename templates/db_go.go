package templates

const DbGo = `package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"

	_ "github.com/$1/$2/config"
)

var DB *gorm.DB = nil

func init() {
	var err error = nil

	connectionStr := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
	)

	// TODO: Figure out how to add defer DB.Close()
	DB, err = gorm.Open("postgres", connectionStr)
	if err != nil {
		panic(fmt.Sprintf("Unable to establish database connection: %v", err))
	}
}
`
