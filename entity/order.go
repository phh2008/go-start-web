package entity

import (
	"com.phh/start-web/util/timeutil"
	"github.com/shopspring/decimal"
)

type Order struct {
	Id       int
	Name     string
	Amount   decimal.Decimal
	CreateAt timeutil.LocalDateTime
	Status   int
	UserId   int
}

func (Order) TableName() string {
	return "order"
}
