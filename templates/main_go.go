package templates

// MainGo is the template for a Gin project's main.go.
const MainGo = `package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/toorop/gin-logrus"

	_ "github.com/$1/$2/config"
	"github.com/$1/$2/controllers"
	_ "github.com/$1/$2/db"
	_ "github.com/$1/$2/docs"
)

// BaseVersion is the SemVer-formatted string that defines the current version of $2.
// Build information will be added at compile-time.
const BaseVersion = "0.1.0-develop"
// BuildTime is a timestamp of when the build is run. This variable is set at compile-time.
var BuildTime string
// GitRevision is the current Git commit ID. If the tree is dirty at compile-time, an "x-" is prepended to the hash.
// This variable is set at compile-time.
var GitRevision string
// GitBranch is the name of the active Git branch at compile-time. This variable is set at compile-time.
var GitBranch string

// @title $1 $2
// @version 0.1.0+0
// @description UPDATE DESCRIPTION FIELD

// @contact.name UPDATE CONTACT NAME
// @contact.email UPDATE CONTACT EMAIL

// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host UPDATE HOST
// @BasePath /api/v1
func main() {
	version := fmt.Sprintf(
		"%s+%s.%s.%s",
		BaseVersion,
		GitBranch,
		GitRevision,
		BuildTime,
	)

	log := logrus.New()
	router := gin.New()
	router.Use(ginlogrus.Logger(log), gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		// Visit {host}/api/v1/swagger/index.html to see the API documentation.
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

        health := controllers.NewHealthController(version)
		{
			v1.GET("/health", health.Check)
		}
	}

	_ = router.Run(viper.GetString("url"))
}
`
