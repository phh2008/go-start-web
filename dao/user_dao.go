package dao

import (
	"com.phh/start-web/entity"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserDaoSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

type UserDao struct {
	Db *gorm.DB
}

func (a *UserDao) GetById(id int) entity.User {
	var user entity.User
	a.Db.Model(&entity.User{}).Take(&user, id)
	return user
}
