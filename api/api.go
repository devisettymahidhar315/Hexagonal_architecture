package api

import (
	_ "app/docs"
	"app/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Your API Title
// @version 1.0
// @description This is a sample server for Your API.
// @host localhost:8080
// @BasePath /
func Router(c *gin.Engine, q *router.Router) {
	// Initializes the swagger
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define routes
	c.GET("/print", q.Print)
	c.POST("/start/:key", q.Insertatstarting)
	c.POST("/end/:key", q.Insertatending)
	c.DELETE("/front", q.DelAtBeg)
	c.DELETE("/end", q.DelAtEnd)
	c.GET("/length", q.Length)
}
