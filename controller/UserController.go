package controller

import (
	"fmt"
	"net/http"

	"github.com/alijkdkar/ArvanChallenge/domain"
	"github.com/alijkdkar/ArvanChallenge/pkg"
	"github.com/alijkdkar/ArvanChallenge/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserHandlerRegister(ctx *gin.Engine) {
	baseUrl := "/user"
	ctx.GET(baseUrl+"/all", getAll)
	ctx.POST(baseUrl+"", createUser)
	ctx.DELETE(baseUrl+"/:id", deleteUser)
	ctx.GET(baseUrl+"/:id", getUserDetail)
	ctx.PUT(baseUrl+"/:id", updateUser)
	ctx.GET(baseUrl+"/discountused/:Id", GetUsersUsedDiscount)

}

func createUser(ctx *gin.Context) {
	rep := repository.NewUserRepository()

	var userReq createUserCommand
	errJson := ctx.ShouldBindJSON(&userReq)

	if errJson != nil {
		fmt.Println("error on Cast create user to command", errJson)
		pkg.BadRequestError(ctx)
		return
	}

	if userReq.MobileNumber == "" {
		ctx.JSON(http.StatusAccepted, gin.H{
			"message": "Mobile number is reqired.",
		})
		return
	}

	user := domain.CreateNewUser(userReq.Name, userReq.LastName, userReq.MobileNumber)
	user.NewInstance()
	err := rep.Create(user)
	if err != nil {
		pkg.ServerSideError(ctx)
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "User Saved",
	})
}

func getAll(ctx *gin.Context) {
	rep := repository.NewUserRepository()
	allUser := rep.GetUsers()
	ctx.JSON(http.StatusAccepted, allUser)
}

func getUserDetail(ctx *gin.Context) {

	id := ctx.Param("id")
	fmt.Println(id)
	if id == "" {
		pkg.BadRequestError(ctx)
		return
	}
	rep := repository.NewUserRepository()
	user, err := rep.GetUserById(uuid.MustParse(id))
	if err != nil {
		pkg.NotFoundError(ctx)
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}

func deleteUser(ctx *gin.Context) {

	id := ctx.Param("id")
	if id == "" {
		pkg.BadRequestError(ctx)
		return
	}
	rep := repository.NewUserRepository()
	err := rep.DeleteById(uuid.MustParse(id))

	if err != nil {
		pkg.ServerSideError(ctx)
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "user deleted",
	})
}

func updateUser(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		fmt.Println("Id is empty")
		pkg.BadRequestError(ctx)
		return
	}

	var userReq createUserCommand
	errJson := ctx.ShouldBindJSON(&userReq)

	if errJson != nil {
		fmt.Println("error on cast command:", errJson)
		pkg.BadRequestError(ctx)
		return
	}

	rep := repository.NewUserRepository()

	user := domain.CreateNewUser(userReq.Name, userReq.LastName, userReq.MobileNumber)
	user.Id = uuid.MustParse(id)
	user.SetVersion(userReq.Version)
	uperr := rep.Update(user)
	if uperr != nil {
		fmt.Println("error on update: ", uperr)
		pkg.DataVersionCustomError(ctx, uperr.Error())
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "user updated",
	})
}

func GetUsersUsedDiscount(ctx *gin.Context) {
	id := ctx.Param("Id")
	if id == "" {
		fmt.Println("Id Got:", id)
		pkg.BadRequestError(ctx)
		return
	}

	_cardRepo := repository.NewCreditCardRepository()
	lis := _cardRepo.GetUsersUsedDiscount(uuid.MustParse(id))

	ctx.JSON(http.StatusOK, lis)
}

type createUserCommand struct {
	Name         string `josn:"name"`
	LastName     string `json:"lastName"`
	MobileNumber string `json:"mobileNumber"`
	Version      uint   `json:"version"`
}
