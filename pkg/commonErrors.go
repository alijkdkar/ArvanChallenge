package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestError(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "400 - bad request",
	})

}

func ServerSideError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "500 - internal server error",
	})
}

func NotFoundError(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"message": "404 - not found",
	})
}

func DataVersionError(ctx *gin.Context) {
	ctx.JSON(http.StatusConflict, gin.H{
		"message": "409 - this error modify by other user",
	})
}
