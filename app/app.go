package app

import (
	"gopkg.in/gin-gonic/gin.v1"
)

var Router *gin.Engine

func init() {
	initLogger()

	Log.Debug("Loading environment vars...")
	initConf()

	Log.Debug("Connecting to database...")
	initDb()

	if Env("ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	Router = gin.Default()
}
