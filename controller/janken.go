package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"math/rand"
)

type jankenGetRes struct {
	Janken string `json:"相手が出した手"`
}

func JankenGetController(c *gin.Context){
	judge := rand.Intn(3)
	janken := "void"
	switch judge {
	case 0:
		janken = "pa"
	case 1:
		janken = "gu"
	case 2:
		janken = "tyoki"
	}
	c.JSON(http.StatusOK, jankenGetRes{
		Janken: janken,
	})
}