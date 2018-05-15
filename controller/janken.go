package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"math/rand"
)

type jankenGetRes struct {
	janken int `json:username`
}

func JankenGetController(c *gin.Context){
	c.JSON(http.StatusOK, jankenGetRes{
		janken: rand.Intn(3),
	})
}