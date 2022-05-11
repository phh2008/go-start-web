package entity

import (
	"com.phh/start-web/util"
	"github.com/shopspring/decimal"
)

type Order struct {
	Id       int
	Name     string
	Amount   decimal.Decimal
	CreateAt util.LocalDateTime
	Status   int
	UserId   int
}

func (Order) TableName() string {
	return "order"
}
