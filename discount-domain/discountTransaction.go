package discountdomain

import (
	"time"

	"github.com/google/uuid"
)

type DiscountTransaction struct {
	Id         uuid.UUID
	DiscountId uuid.UUID
	CardId     uuid.UUID
	CreatedAt  time.Time
}

func CreateNewDiscountTransactionInstance(discountId uuid.UUID, cardId uuid.UUID) *DiscountTransaction {

	return &DiscountTransaction{
		Id:         uuid.New(),
		DiscountId: discountId,
		CardId:     cardId,
		CreatedAt:  time.Now(),
	}
}

func (tr *DiscountTransaction) DiscountTranssToMap() map[string]interface{} {
	DiscountTarMap := make(map[string]interface{})
	DiscountTarMap["DiscountId"] = tr.DiscountId
	DiscountTarMap["CardId"] = tr.CardId
	DiscountTarMap["CreatedAt"] = tr.CreatedAt
	DiscountTarMap["Id"] = tr.Id
	return DiscountTarMap
}
