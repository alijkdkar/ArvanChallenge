package domain

import "github.com/google/uuid"

type Transaction struct {
	BaseEntity
	Amount          float64
	Type            uint
	IsDiscount      bool
	CardNumberRefer uuid.UUID
	DiscountReferId uuid.UUID
}

func CreateNewTransaction(amount float64, typ uint, cardId uuid.UUID, isDiscount bool, discountReferId uuid.UUID) *Transaction {
	newITem := Transaction{
		Amount:          amount,
		Type:            typ,
		IsDiscount:      isDiscount,
		CardNumberRefer: cardId,
		DiscountReferId: discountReferId,
	}
	newITem.NewInstance()
	return &newITem
}
