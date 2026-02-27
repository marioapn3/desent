package handler

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func Echo(c *gin.Context) {
	var body map[string]interface{}
	c.BindJSON(&body)
	c.JSON(200, body)
}
