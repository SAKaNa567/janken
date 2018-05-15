package main

import (
	"github.com/gin-gonic/gin"
	"Web_Application_By_Go/voyage/gin-api/controller"
)

func main() {
	m := gin.Default()
	m.POST("/game",controller.UserPostController)
	m.Run(":8080")
}
