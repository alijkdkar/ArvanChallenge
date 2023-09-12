package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterControllers(ctx *gin.Engine) {
	UserHandlerRegister(ctx)
	//Add Other Controllers
}
