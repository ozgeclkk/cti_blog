package main

import (
	"blog/config"
	"blog/database"
	"blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.InitDB()
	database.SeedAdmin()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(config.SunucuPort)

}
