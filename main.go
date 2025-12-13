package main

import (
	"hltv/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")
	controllers.HLTVendpoints(api)
	router.Run(":8080")
}
