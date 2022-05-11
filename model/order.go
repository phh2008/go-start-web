package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderQuery struct {
	Page
	Status int
	Name   string
	UserId int
	Id     int
}

type OrderResult struct {
	Id       int
	Name     string
	UserId   int
	Amount   decimal.Decimal
	CreateAt time.Time
	Status   int
}
