package sysdao

import (
	"com.phh/start-web/entity/sysentity"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserDaoSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

type UserDao struct {
	Db *gorm.DB
}

func (a *UserDao) GetById(id int) sysentity.User {
	var user sysentity.User
	a.Db.Model(&sysentity.User{}).Take(&user, id)
	return user
}
