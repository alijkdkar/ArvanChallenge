package discountcontroller

import (
	"fmt"
	"net/http"
	"time"

	discountdomain "github.com/alijkdkar/ArvanChallenge/discount-domain"
	discountrepository "github.com/alijkdkar/ArvanChallenge/discount-repository"
	"github.com/alijkdkar/ArvanChallenge/pkg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterDiscountCountEndPoint(ctx *gin.Engine) {
	baseUrl := "/discount"
	ctx.POST(baseUrl+"", CreateDiscountOppurtunity)
	ctx.DELETE(baseUrl+"/:Id", DeleteDiscount)
	ctx.POST(baseUrl+"/:Id/use/:cardId", UseOfDiscount)
}

func CreateDiscountOppurtunity(ctx *gin.Context) {
	var request CreateDiscountDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println("cast create Discount command err:", err)
		pkg.BadRequestError(ctx)
		return
	}
	_dicRepo := discountrepository.CreateNewDiscountRepositoryInstance()

	if request.EnableTime == "" {
		newDisc := *discountdomain.CreateNewDiscountOpportunity(request.Name, request.MaxCount, request.Amount, time.Now(), request.Code)
		err := _dicRepo.CreateDiscountOpp(newDisc)
		if err != nil {
			pkg.ServerSideError(ctx)
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": "Discount Created Successfuly and usable after enbale time. if you set 'EnableTime' empty its can use from right now.", "Id": newDisc.Id.String()})
	} else {
		discountDuoTime, er := time.Parse("2006-01-02T15:04:05.000Z", request.EnableTime)
		if er != nil {
			pkg.BadRequestError(ctx)
			return
		}
		err := _dicRepo.CreateDiscountOpp(*discountdomain.CreateNewDiscountOpportunity(request.Name, request.MaxCount, request.Amount, discountDuoTime, request.Code))
		if err != nil {
			pkg.ServerSideError(ctx)
			return
		}
	}

}

func DeleteDiscount(ctx *gin.Context) {
	Id := ctx.Param("Id")
	if Id == "" {
		pkg.BadRequestError(ctx)
		return
	}
	_repoDisc := discountrepository.CreateNewDiscountRepositoryInstance()

	if !_repoDisc.ExistsDiscountOpp(uuid.MustParse(Id)) {
		pkg.NotFoundError(ctx)
		return
	}

	err := _repoDisc.RemoveDiscountOpp(uuid.MustParse(Id))
	if err != nil {
		fmt.Println(err)
		pkg.ServerSideError(ctx)
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Discount Opportunity Deleted"})

}

func UseOfDiscount(ctx *gin.Context) {
	Id := ctx.Param("Id")
	cardId := ctx.Param("cardId")
	if Id == "" || cardId == "" {
		fmt.Println("Id or Card Id is empty")
		pkg.BadRequestError(ctx)
		return
	}

	_dicRepo := discountrepository.CreateNewDiscountRepositoryInstance()
	if !_dicRepo.ExistsDiscountOpp(uuid.MustParse(Id)) {
		pkg.NotFoundError(ctx)
		return
	}

	newUse := discountdomain.CreateNewDiscountTransactionInstance(uuid.MustParse(Id), uuid.MustParse(cardId))

	saveErr := _dicRepo.AddNewDiscountUse(newUse)
	if saveErr != nil {
		pkg.CommonError(ctx, http.StatusOK, saveErr.Error())
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "your Discount Saved"})

}

type CreateDiscountDto struct {
	Name       string `json:"Name"`
	MaxCount   int    `json:"MaxCount"`
	Amount     int64  `json:"Amount"`
	EnableTime string `json:"EnableTime"`
	Code       string `json:"Code"`
}
