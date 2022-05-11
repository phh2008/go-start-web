package service

import (
	"com.phh/start-web/dao"
	"com.phh/start-web/entity"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserService), "*"))

type UserService struct {
	UserDao *dao.UserDao
}

func (a *UserService) GetById(id int) entity.User {
	return a.UserDao.GetById(id)
}
