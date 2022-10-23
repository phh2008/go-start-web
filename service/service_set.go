package service

import (
	"com.phh/start-web/service/odrservice"
	"com.phh/start-web/service/sysservice"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	sysservice.UserSet,
	odrservice.OrderSet,
)
