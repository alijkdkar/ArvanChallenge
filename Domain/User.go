package domain

import "github.com/google/uuid"

type User struct {
	BaseEntity
	MobileNumber string       `gorm:"index"`
	FirstName    string       `json:"FirstName"`
	LastName     string       `json:"LastName"`
	CreditCards  []CreditCard `gorm:"foreignKey:UserRefer"`
}

func CreateNewUser(firsName, lastName, mobileNumber string) *User {
	newUser := &User{FirstName: firsName, LastName: lastName, MobileNumber: mobileNumber}
	newUser.NewInstance()
	return newUser
}

func (us *User) SetVersion(version uint) {
	us.Version = version
}

type IUserRepository interface {
	Create(user *User) error
	GetUserById(id uuid.UUID) (User, error)
	Update(user *User) error
	DeleteById(id uuid.UUID) error
	GetUsers() []User
}
