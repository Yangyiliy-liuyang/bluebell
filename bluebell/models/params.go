package models

type ParamSignUp struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
