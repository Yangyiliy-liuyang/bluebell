package routers

import (
	"bluebell/controller"
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() (r *gin.Engine) {
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/signup", controller.SignUpHander)
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return
}
