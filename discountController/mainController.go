package discountcontroller

import "github.com/gin-gonic/gin"

// register all endpoint
func RegisterDiscountServices(ctx *gin.Engine) {
	RegisterDiscountCountEndPoint(ctx)
	EnableMessageBusServices()
}
