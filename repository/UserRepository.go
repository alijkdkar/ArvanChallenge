package repository

import (
	"github.com/alijkdkar/ArvanChallenge/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Db: *DB,
	}
}

// check UserRepository Can satisfy
var _ domain.IUserRepository = (*UserRepository)(nil)

func (rp *UserRepository) Create(user *domain.User) error {
	res := rp.Db.Create(user)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (rp *UserRepository) GetUserById(id uuid.UUID) domain.User {
	result := domain.User{}

	rp.Db.First(&result, "id = ?", id)
	return result
}

func (rp *UserRepository) GetUsers() []domain.User {
	result := []domain.User{}

	rp.Db.Find(&result)
	return result
}

func (rp *UserRepository) Update(user *domain.User) error {

	dbData := domain.User{}
	rp.Db.First(dbData)

	err := user.UpdateInstance(dbData.Version)
	if err != nil {
		return err
	}
	return rp.Db.Save(user).Error
}

func (rp *UserRepository) DeleteById(id uuid.UUID) error {
	res := rp.Db.Delete(domain.User{}, id)
	return res.Error
}
