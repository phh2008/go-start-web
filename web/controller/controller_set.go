package controller

import "github.com/google/wire"

var ControllerSet = wire.NewSet(
	UserSet,
	OrderSet,
	AuthSet,
)
