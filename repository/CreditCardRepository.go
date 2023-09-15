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

func (rp *CreditCardRepository) GetUsersUsedDiscount(discountId uuid.UUID) []*domain.User {
	users := []*domain.User{}

	transactions := []domain.Transaction{}
	rp.Db.Find(&transactions, "discount_refer_id = ?", discountId)
	cardIds := []uuid.UUID{}
	for _, v := range transactions {
		cardIds = append(cardIds, v.CardNumberRefer)
	}

	cards := []*domain.CreditCard{}
	rp.Db.Where("id IN ?", cardIds).Find(&cards)
	userIds := []uuid.UUID{}
	for _, v := range cards {
		userIds = append(userIds, v.UserRefer)
	}

	rp.Db.Where("id IN ?", userIds).Find(&users)
	return users
}

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

func (rp *CreditCardRepository) AddTransaction(transaction *domain.Transaction) error {

	return rp.Db.Transaction(func(tx *gorm.DB) error {

		var card domain.CreditCard
		rp.Db.First(&card, "id = ?", transaction.CardNumberRefer)
		if card.Id == uuid.Nil {
			return errors.New("card not found")
		}

		if transaction.Type == 1 {
			if err := card.SetAmount(card.Amount + (transaction.Amount * 1)); err != nil {
				return err
			}
		} else if transaction.Type == 2 {
			if err := card.SetAmount(card.Amount + (transaction.Amount * -1)); err != nil {
				return err
			}
		} else {
			return errors.New("server side error ")
		}

		rp.Db.Save(&card)

		if newTran := rp.Db.Create(transaction); newTran.Error != nil {
			return newTran.Error
		}

		return nil
	})

}

func (rp *CreditCardRepository) GetTransactions(cardId uuid.UUID) []*domain.Transaction {
	result := []*domain.Transaction{}
	rp.Db.Find(&result, "card_number_refer = ?", cardId)
	return result
}
