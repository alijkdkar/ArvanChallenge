package domain

import "github.com/google/uuid"

type Transaction struct {
	BaseEntity
	Amount          float64
	Type            uint
	IsBounce        bool
	CardNumberRefer uuid.UUID
}

func CreateNewTransaction(amount float64, typ uint, cardId uuid.UUID, isBounce bool) *Transaction {
	newITem := Transaction{
		Amount:          amount,
		Type:            typ,
		IsBounce:        isBounce,
		CardNumberRefer: cardId,
	}
	newITem.NewInstance()
	return &newITem
}
