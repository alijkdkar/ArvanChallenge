package controller

import (
	"github.com/gin-gonic/gin"
)

// register end point to gin or enable message bus agent
func RegisterControllers(ctx *gin.Engine) {

	ctx.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	UserHandlerRegister(ctx)
	CreditHandlerRegisters(ctx)
	EnableCoreMessageBusServices()
	//Add Other Controllers
}
