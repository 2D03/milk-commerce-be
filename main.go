package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	env := os.Getenv("env")
	version := os.Getenv("version")
	config := os.Getenv("config")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"env":     env,
			"config":  config,
			"version": version,
		})
	})

	_ = router.Run(":8000")
}
