package discountdomain

import (
	"time"

	"github.com/google/uuid"
)

type DiscountOpportunity struct {
	Id         uuid.UUID
	Name       string
	Code       string
	MaxCount   int
	Amount     int64
	EnableTime time.Time
	UsedCount  uint
}

func (th *DiscountOpportunity) SetId(id uuid.UUID) {
	th.Id = id
}

func (user DiscountOpportunity) DiscountToMap() map[string]interface{} {
	DiscountMap := make(map[string]interface{})
	DiscountMap["Id"] = user.Id
	DiscountMap["Name"] = user.Name
	DiscountMap["Code"] = user.Code
	DiscountMap["MaxCount"] = user.MaxCount
	DiscountMap["EnableTime"] = user.EnableTime
	DiscountMap["UsedCount"] = user.UsedCount
	return DiscountMap
}

func CreateNewDiscountOpportunity(name string, maxCount int, amount int64, enableTime time.Time, code string) *DiscountOpportunity {

	return &DiscountOpportunity{
		Id:         uuid.New(),
		Name:       name,
		MaxCount:   maxCount,
		Amount:     amount,
		EnableTime: enableTime,
		UsedCount:  0,
		Code:       code,
	}
}

type IDiscountOpportunityRepository interface {
	CreateDiscountOpp(disc DiscountOpportunity) error
	RemoveDiscountOpp(Id uuid.UUID) error
	ExistsDiscountOpp(Id uuid.UUID) bool
}
