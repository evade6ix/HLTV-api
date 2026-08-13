package main

import (
	"embed"
	"io/fs"
	"net/http"

	"hltv/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "hltv/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed dashboard/*
var dashboard embed.FS

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	controllers.HLTVendpoints(api)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	dashboardFS, _ := fs.Sub(dashboard, "dashboard")
	router.StaticFS("/assets", http.FS(dashboardFS))
	router.GET("/", func(c *gin.Context) {
		c.FileFromFS("index.html", http.FS(dashboardFS))
	})
	router.Run(":8080")
}
