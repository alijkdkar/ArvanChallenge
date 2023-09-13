package pkg

import (
	"fmt"
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

func DataVersionCustomError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusConflict, gin.H{
		"message": "409 - this error modify by other user",
	})
}

func CommonError(ctx *gin.Context, httpStatusCode int, message string) {
	ctx.JSON(httpStatusCode, gin.H{
		"message": fmt.Sprintf("%v - %v", httpStatusCode, message),
	})
}
