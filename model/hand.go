package model

type UserPostReq struct {
	Userhand string `json:"自分が出した手"`
	Computerhand string `json:"コンピュータが出した手"`
	Which string `json:"勝者"`
}