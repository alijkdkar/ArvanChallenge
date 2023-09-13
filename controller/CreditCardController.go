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

func CreditHandlerRegisters(ctx *gin.Engine) {
	baseUrl := "/card"
	ctx.GET(baseUrl+"/all/:Id", getAllCards)
	ctx.POST(baseUrl+"/:Id", createCard)
	ctx.GET(baseUrl+"/:userId/detail/:Id", getCardDetail)
	ctx.DELETE(baseUrl+"/:Id", deleteCard)
	ctx.PUT(baseUrl+"/:userId/update/:Id", updateCard)
}

func getAllCards(ctx *gin.Context) {
	id := ctx.Param("Id")
	if id == "" {
		pkg.BadRequestError(ctx)
		return
	}
	rep := repository.NewCreditCardRepository()
	cards := rep.GetCreditCards(uuid.MustParse(id))
	ctx.JSON(http.StatusOK, cards)
}

func createCard(ctx *gin.Context) {

	userId := ctx.Param("Id")

	if userId == "" {
		pkg.BadRequestError(ctx)
		return
	}

	//check user exits
	_userRepo := repository.NewUserRepository()
	if _, e := _userRepo.GetUserById(uuid.MustParse(userId)); e != nil {
		pkg.NotFoundError(ctx)
		return
	}

	var request CardDto
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		fmt.Println("error on requst Command Cast:", err)
		pkg.ServerSideError(ctx)
		return
	}
	carRepo := repository.NewCreditCardRepository()
	newCard := domain.CreateCreditCardNewInstance(request.CardNumber, uuid.MustParse(userId))
	errCreate := carRepo.Create(newCard)
	if errCreate != nil {
		fmt.Println("eeror on create Card : ", errCreate)
		pkg.ServerSideError(ctx)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Card Created"})
}

func getCardDetail(ctx *gin.Context) {
	userId := ctx.Param("userId")
	cardId := ctx.Param("Id")

	if userId == "" || cardId == "" {
		pkg.BadRequestError(ctx)
		return
	}

	//check user exits
	_userRepo := repository.NewUserRepository()
	if _, e := _userRepo.GetUserById(uuid.MustParse(userId)); e != nil {
		pkg.NotFoundError(ctx)
		return
	}

	_cardRep := repository.NewCreditCardRepository()
	card, err := _cardRep.GetById(uuid.MustParse(cardId))
	if err != nil {
		fmt.Println("get card detail:", err)
		pkg.ServerSideError(ctx)
		return
	}
	ctx.JSON(http.StatusOK, card)
}

func deleteCard(ctx *gin.Context) {
	id := ctx.Param("Id")
	if id == "" {
		pkg.BadRequestError(ctx)
	}
	_cardRep := repository.NewCreditCardRepository()
	if err := _cardRep.DeleteById(uuid.MustParse(id)); err != nil {
		fmt.Println("delete card error : ", err)
		pkg.ServerSideError(ctx)
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "card Deleted"})

}

func updateCard(ctx *gin.Context) {
	userId := ctx.Param("userId")
	cardId := ctx.Param("Id")

	if userId == "" || cardId == "" {
		pkg.BadRequestError(ctx)
		return
	}

	//check user exits
	_userRepo := repository.NewUserRepository()
	if _, e := _userRepo.GetUserById(uuid.MustParse(userId)); e != nil {
		pkg.NotFoundError(ctx)
		return
	}
	var request domain.CreditCard
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println("command casting error:", err)
		pkg.ServerSideError(ctx)
		return
	}

	_cardRep := repository.NewCreditCardRepository()
	cardUpdate := domain.CreateCreditCardNewInstance(request.CardNumber, uuid.MustParse(userId))

	cardUpdate.SetVersion(request.Version)
	cardUpdate.SetId(uuid.MustParse(cardId))

	if err := _cardRep.Update(cardUpdate); err != nil {
		fmt.Println("card update error:", err)
		pkg.CommonError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

}

type CardDto struct {
	CardNumber string `json:"CardNumber"`
	Version    uint   `json:"Version"`
	// UserId     uuid.UUID `json:"UserId"`
}
