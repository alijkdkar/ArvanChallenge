package domain

type User struct {
	BaseEntity
	MobileNumber string `gorm:"index"`
	FirstName    string
	LastName     string
	CreditCards  []CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
