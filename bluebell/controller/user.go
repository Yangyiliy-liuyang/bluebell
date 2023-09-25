package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// 处理注册请求的函数
func SignUpHander(ctx *gin.Context) {
	//1.获取参数和参数校验
	var p models.ParamSignUp
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invilid param", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"code": 201,
			"msg":  "请求参数有误",
		})
		return
	}
	fmt.Println(p)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	//2.业务处理
	logic.SignUp()

	//3.返回响应
	ctx.JSON(http.StatusOK, "")
}
