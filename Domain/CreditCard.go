package domain

import (
	"errors"

	"github.com/google/uuid"
)

type CreditCard struct {
	BaseEntity
	CardNumber   string
	Amount       float64
	Transactions []*Transaction `gorm:"foreignKey:CardNumberRefer"`
	UserRefer    uuid.UUID
}

func CreateCreditCardNewInstance(cardNumber string, userId uuid.UUID) *CreditCard {
	card := &CreditCard{
		CardNumber: cardNumber,
		UserRefer:  userId,
	}
	card.NewInstance()
	return card

}

func (card *CreditCard) SetVersion(version uint) {
	card.Version = version
}
func (card *CreditCard) SetId(id uuid.UUID) {
	card.Id = id
}
func (card *CreditCard) SetAmount(amount float64) error {

	if amount < 0 {
		return errors.New("card amount can not be negative")
	}
	card.Amount = amount
	return nil
}

type ICreditCardRepositroy interface {
	Create(card *CreditCard) error
	GetById(id uuid.UUID) (CreditCard, error)
	Update(card *CreditCard) error
	DeleteById(id uuid.UUID) error
	GetCreditCards(id uuid.UUID) []CreditCard
	AddTransaction(transaction *Transaction) error
	GetTransactions(id uuid.UUID) []*Transaction
}
