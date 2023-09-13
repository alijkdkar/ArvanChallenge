package domain

import "github.com/google/uuid"

type CreditCard struct {
	BaseEntity
	CardNumber   string
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

type ICreditCardRepositroy interface {
	Create(card *CreditCard) error
	GetById(id uuid.UUID) (CreditCard, error)
	Update(card *CreditCard) error
	DeleteById(id uuid.UUID) error
	GetCreditCards(id uuid.UUID) []CreditCard
	AddTransaction(transaction *Transaction) error
}
