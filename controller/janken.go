package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Web_Application_By_Go/voyage/gin-api/model"
	"Web_Application_By_Go/voyage/gin-api/service"
)

func UserPostController(c *gin.Context) {
	json := &model.UserPostReq{}
	service.ComputerHand(json)//コンピュタの手を生成+格納
	c.ShouldBindJSON(json)//ユーザの手を格納
	service.JudgeWinner(json)
	c.JSON(http.StatusOK,json)
}