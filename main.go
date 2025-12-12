package main

import (
  "shotner/database"

	"github.com/gin-gonic/gin"
)

func main() {
  database.OpenConnection()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
  router.Run()
}