package service

import (
	"math/rand"
	"Web_Application_By_Go/voyage/gin-api/model"
)

func ComputerHand(req *model.UserPostReq) {
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
	req.Computerhand = janken
}

func JudgeWinner(req *model.UserPostReq) {
	if req.Userhand == "tyoki" && req.Computerhand == "pa"{
		req.Which = "user"
	}
}