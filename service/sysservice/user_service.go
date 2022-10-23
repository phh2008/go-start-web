package sysservice

import (
	"com.phh/start-web/dao/sysdao"
	entity "com.phh/start-web/entity/sysentity"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserService), "*"))

type UserService struct {
	UserDao *sysdao.UserDao
}

func (a *UserService) GetById(id int) entity.User {
	return a.UserDao.GetById(id)
}
