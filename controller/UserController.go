package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	repository "github.com/alijkdkar/ArvanChallenge/Repository"
	"github.com/alijkdkar/ArvanChallenge/domain"
	"github.com/alijkdkar/ArvanChallenge/pkg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserHandlerRegister(ctx *gin.Engine) {
	baseUrl := "/user"
	ctx.GET(baseUrl+"/all", getAll)
	ctx.POST(baseUrl, createUser)
	ctx.DELETE(baseUrl+"", deleteUser)
	ctx.GET(baseUrl+"/:id", getUserDetail)
	ctx.PUT(baseUrl+"", updateUser)

}

func User(ctx *gin.Context) {
	//Create one  and Get All User
	rep := repository.NewUserRepository()

	decoder := json.NewDecoder(ctx.Request.Body)
	var userReq createUserCommand
	errJson := decoder.Decode(&userReq)

	if errJson != nil {
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
	err := rep.Create(user)
	if err != nil {
		pkg.ServerSideError(ctx)
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "User Saved",
	})

}

func createUser(ctx *gin.Context) {
	rep := repository.NewUserRepository()

	decoder := json.NewDecoder(ctx.Request.Body)
	var userReq createUserCommand
	errJson := decoder.Decode(&userReq)

	if errJson != nil {
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

	fmt.Println("in get detail ", ctx.Request.URL.Path)
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
	userJson, errr := json.Marshal(user)
	if errr != nil {
		pkg.BadRequestError(ctx)
		return
	}
	ctx.JSON(http.StatusAccepted, userJson)
}

func deleteUser(ctx *gin.Context) {

	id := strings.TrimPrefix(ctx.Request.URL.Path, "/user/delete/")
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
	id := strings.TrimPrefix(ctx.Request.URL.Path, "/user/delete/")
	if id == "" {
		pkg.BadRequestError(ctx)
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var userReq createUserCommand
	errJson := decoder.Decode(&userReq)

	if errJson != nil {
		pkg.BadRequestError(ctx)
		return
	}

	rep := repository.NewUserRepository()
	userDb, err := rep.GetUserById(uuid.MustParse(id))

	if err != nil {
		pkg.NotFoundError(ctx)
		return
	}
	if err := userDb.UpdateInstance(userReq.Version); err != nil {
		pkg.DataVersionError(ctx)
		return
	}

	uperr := rep.Update(&userDb)
	if uperr != nil {
		pkg.ServerSideError(ctx)
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "user updated",
	})
}

type createUserCommand struct {
	Name         string `josn:"name"`
	LastName     string `json:"lastName"`
	MobileNumber string `json:"mobileNumber"`
	Version      uint   `json:"version"`
}
