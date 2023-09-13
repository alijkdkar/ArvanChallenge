package repository

import (
	"errors"
	"fmt"

	"github.com/alijkdkar/ArvanChallenge/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreditCardRepository struct {
	Db gorm.DB
}

func NewCreditCardRepository() *CreditCardRepository {
	return &CreditCardRepository{
		Db: *DB,
	}
}

var _ domain.ICreditCardRepositroy = (*CreditCardRepository)(nil)

func (rp *CreditCardRepository) Create(card *domain.CreditCard) error {
	result := rp.Db.Create(card)
	return result.Error
}

func (rp *CreditCardRepository) GetById(id uuid.UUID) (domain.CreditCard, error) {
	card := domain.CreditCard{}

	rp.Db.First(&card, "id = ?", id)
	if card.Id == uuid.Nil {
		return card, errors.New("card not found")
	}
	return card, nil
}

func (rp *CreditCardRepository) Update(card *domain.CreditCard) error {

	dbData := domain.CreditCard{}
	rp.Db.First(&dbData, "id = ?", card.Id)

	fmt.Println(dbData)

	if dbData.Id == uuid.Nil {
		return errors.New("card not found")
	}

	err := card.UpdateInstance(dbData.Version)
	if err != nil {
		return err
	}

	if exResult := rp.Db.Save(card); exResult.Error != nil {
		return errors.New("server side error")
	}
	return nil
}

func (rp *CreditCardRepository) DeleteById(id uuid.UUID) error {
	res := rp.Db.Delete(domain.CreditCard{}, "id = ?", id)
	return res.Error
}

func (rp *CreditCardRepository) GetCreditCards(UserId uuid.UUID) []domain.CreditCard {
	cards := []domain.CreditCard{}
	rp.Db.Find(&cards, "user_refer = ?", UserId)
	return cards
}
