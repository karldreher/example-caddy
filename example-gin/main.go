package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthz", GetHealthz)
	router.GET("/api/v1/time", apiv1GetTime)
	router.Run(":8080")
}

func GetHealthz(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func apiv1GetTime(c *gin.Context) {
	t := time.Now()
	c.JSON(200, gin.H{"time": t})
}
