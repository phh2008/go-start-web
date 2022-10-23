package odrdao

import (
	"com.phh/start-web/entity/odrentity"
	"com.phh/start-web/model"
	"com.phh/start-web/util"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var OrderSet = wire.NewSet(wire.Struct(new(OrderDAO), "*"))

type OrderDAO struct {
	Db *gorm.DB
}

func (a *OrderDAO) GetById(id int) odrentity.Order {
	var order odrentity.Order
	a.Db.Model(&odrentity.Order{}).Take(&order, id)
	return order
}

func (a *OrderDAO) ListByUserId(userId int) []odrentity.Order {
	var orders []odrentity.Order
	a.Db.Model(&odrentity.Order{}).Where("user_id=?", userId).Find(&orders)
	return orders
}

func (a *OrderDAO) Query(orderQuery model.OrderQuery) model.Page {
	db := a.Db
	db = db.Model(&odrentity.Order{})
	if orderQuery.Id != 0 {
		db = db.Where("id=?", orderQuery.Id)
	}
	if orderQuery.Name != "" {
		db = db.Where("name like ?", "%"+orderQuery.Name+"%")
	}
	if orderQuery.UserId != 0 {
		db = db.Where("user_id=?", orderQuery.UserId)
	}
	if orderQuery.Status != 0 {
		db = db.Where("status=?", orderQuery.Status)
	}
	var result []model.OrderResult
	var count int64
	var page model.Page
	// 工具深拷贝
	// copier.Copy(&page, &orderQuery.Page)
	// 浅拷贝
	page = orderQuery.Page
	db.Count(&count).Scopes(util.Paginate(page)).Find(&result)
	page.Count = count
	page.Data = result
	return page
}
