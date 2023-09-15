package discountcontroller

import "github.com/gin-gonic/gin"

func RegisterDiscountServices(ctx *gin.Engine) {
	RegisterDiscountCountEndPoint(ctx)
	EnableMessageBusServices()
}
