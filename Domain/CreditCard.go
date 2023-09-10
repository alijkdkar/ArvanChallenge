package domain

import "github.com/google/uuid"

type CreditCard struct {
	BaseEntity
	CardNumber   string
	Transactions []*Transaction `gorm:"foreignKey:CardNumberRefer"`
	UserRefer    uuid.UUID
}
