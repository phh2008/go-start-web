package odrservice

import (
	"com.phh/start-web/dao/odrdao"
	"com.phh/start-web/entity/odrentity"
	"com.phh/start-web/model"
	"com.phh/start-web/model/result"
	"com.phh/start-web/pkg/global"
	"github.com/google/wire"
)

var OrderSet = wire.NewSet(wire.Struct(new(OrderService), "*"))

type OrderService struct {
	OrderDao *odrdao.OrderDAO
}

func (a *OrderService) GetById(id int) odrentity.Order {
	return a.OrderDao.GetById(id)
}

func (a *OrderService) ListByUserId(userId int) []odrentity.Order {
	return a.OrderDao.ListByUserId(userId)
}

func (a *OrderService) Query(orderQuery model.OrderQuery) result.Result {
	page := a.OrderDao.Query(orderQuery)
	global.Log.Infof("%#v", page)
	return result.Success().SetData(page)
}
