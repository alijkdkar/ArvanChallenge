package discountdomain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type TransactionStatus int

const (
	None      TransactionStatus = 0
	Saved     TransactionStatus = 1
	Published TransactionStatus = 2
	Complited TransactionStatus = 3
)

type DiscountTransaction struct {
	Id         uuid.UUID
	DiscountId uuid.UUID
	CardId     uuid.UUID
	CreatedAt  time.Time
	Status     TransactionStatus
	Amount     int64
}

func CreateNewDiscountTransactionInstance(discountId uuid.UUID, cardId uuid.UUID) *DiscountTransaction {

	return &DiscountTransaction{
		Id:         uuid.New(),
		DiscountId: discountId,
		CardId:     cardId,
		CreatedAt:  time.Now(),
		Status:     None,
	}
}

func (tr *DiscountTransaction) ChangeStatus(status TransactionStatus) {
	//can check change status logic
	tr.Status = status
}

func (tr *DiscountTransaction) DiscountTranssToMap() map[string]interface{} {
	DiscountTarMap := make(map[string]interface{})
	DiscountTarMap["DiscountId"] = tr.DiscountId.String()
	DiscountTarMap["CardId"] = tr.CardId.String()
	DiscountTarMap["CreatedAt"] = tr.CreatedAt
	DiscountTarMap["Id"] = tr.Id.String()
	DiscountTarMap["Status"] = fmt.Sprintf("%d", tr.Status)
	DiscountTarMap["Amount"] = fmt.Sprintf("%d", tr.Amount)
	return DiscountTarMap
}

func (tr *DiscountTransaction) LoadFromMap(mp map[string]string) {
	tr.DiscountId = uuid.MustParse(mp["DiscountId"])
	tr.CardId = uuid.MustParse(mp["CardId"])
	created, er := time.Parse("2006-01-02T15:04:05.000Z", mp["CreatedAt"])
	if er == nil {
		tr.CreatedAt = created
	}
	tr.Id = uuid.MustParse(mp["Id"])
	vl, er := strconv.Atoi(mp["Status"])
	if er == nil {
		tr.Status = TransactionStatus(vl)
	}
	am, erAm := strconv.Atoi(mp["Amount"])
	if erAm == nil {
		tr.Amount = int64(am)
	}
}
