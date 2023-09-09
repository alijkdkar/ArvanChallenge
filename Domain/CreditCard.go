package domain

type CreditCard struct {
	BaseEntity
	CardNumber   string
	Transactions []Transaction `gorm:"many2many:user_transactions;"`
}
