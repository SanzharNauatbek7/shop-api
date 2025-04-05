package main

import (
	"github.com/gin-gonic/gin"
	"shop-api/config"
	"shop-api/routes"
)

func main() {
	config.InitMongo() // заменили InitDB()
	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
