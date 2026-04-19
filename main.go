package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luckyparakh/agentclinic/web/pages"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
		pages.Home().Render(c.Request.Context(), c.Writer)
	})

	r.Static("/static", "./static")

	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	setupRouter().Run(":" + port)
}
