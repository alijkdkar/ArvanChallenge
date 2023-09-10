package domain

import "github.com/google/uuid"

type Transaction struct {
	BaseEntity
	CreditCardId    uuid.UUID
	Amount          float64
	Type            uint
	IsBounce        bool
	CardNumberRefer uuid.UUID
}
