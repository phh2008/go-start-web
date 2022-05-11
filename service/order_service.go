package service

import (
	"com.phh/start-web/dao"
	"com.phh/start-web/entity"
	"com.phh/start-web/model"
	"github.com/google/wire"
)

var OrderSet = wire.NewSet(wire.Struct(new(OrderService), "*"))

type OrderService struct {
	OrderDao *dao.OrderDAO
}

func (a *OrderService) GetById(id int) entity.Order {
	return a.OrderDao.GetById(id)
}

func (a *OrderService) ListByUserId(userId int) []entity.Order {
	return a.OrderDao.ListByUserId(userId)
}

func (a *OrderService) Query(orderQuery model.OrderQuery) model.Result {
	page := a.OrderDao.Query(orderQuery)
	return model.Success().SetData(page)
}
