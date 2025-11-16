package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })

	router.POST("/ping",upload_post)

  })
  router.Run() // listens on 0.0.0.0:8080 by default
}

func upload_post(c *gin.Context){}