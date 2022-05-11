package controller

import (
	"com.phh/start-web/model"
	"com.phh/start-web/service"
	"com.phh/start-web/util"
	"fmt"
	"github.com/google/wire"
	"github.com/kataras/iris/v12"
)

var OrderSet = wire.NewSet(wire.Struct(new(OrderController), "*"))

type OrderController struct {
	OrderService *service.OrderService
}

func (a *OrderController) GetById(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	order := a.OrderService.GetById(id)
	ctx.JSON(order)
}

func (a *OrderController) Query(ctx iris.Context) {
	user, _ := ctx.Values().Get("user").(util.UserClaims)
	//TODO 拿到当前登录用户，可把 ctx 传递到 service层
	fmt.Printf("%#v", user)
	var orderQuery model.OrderQuery
	if err := ctx.ReadQuery(&orderQuery); err != nil {
		ctx.JSON(model.NewResult("1000", "参数错误", nil))
		return
	}
	result := a.OrderService.Query(orderQuery)
	ctx.JSON(result)
}
