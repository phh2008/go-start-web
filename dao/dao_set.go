package dao

import (
	"com.phh/start-web/dao/odrdao"
	"com.phh/start-web/dao/sysdao"
	"github.com/google/wire"
)

var DaoSet = wire.NewSet(
	sysdao.UserDaoSet,
	odrdao.OrderSet,
)
