package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	Id        uuid.UUID `gorm:"primaryKey" json:"ids"`
	CreatedBy uuid.UUID
	CreateAt  time.Time
	UpdateBy  uuid.UUID
	UpdateAt  time.Time
	Version   uint
}

func (ent *BaseEntity) NewInstance() {

	ent.Id = uuid.New()
	ent.Version = uint(rand.Int())
	ent.CreateAt = time.Now()
}

func (ent *BaseEntity) UpdateInstance(version uint) error {

	if ent.Version == version {
		ent.Version = uint(rand.Int())
		ent.UpdateAt = time.Now()
		return nil
	}
	return errors.New("this data have edited before")
}
