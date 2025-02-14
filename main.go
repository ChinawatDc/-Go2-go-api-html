package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check route
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "API Gateway is running",
		})
	})

	r.GET("/api/v1/", func(c *gin.Context) {
		content, err := ioutil.ReadFile("./public/html/WelcomeMessage.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Error reading file")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", content)
	})

	r.Run(":8080")
}
