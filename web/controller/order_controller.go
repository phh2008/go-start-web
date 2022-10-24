package controller

import (
	"com.phh/start-web/model"
	"com.phh/start-web/model/result"
	"com.phh/start-web/pkg/common"
	"com.phh/start-web/pkg/global"
	"com.phh/start-web/service/odrservice"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
	"strconv"
)

var OrderSet = wire.NewSet(wire.Struct(new(OrderController), "*"))

type OrderController struct {
	OrderService *odrservice.OrderService
}

func (a *OrderController) GetById(ctx *gin.Context) {
	id := ctx.Query("id")
	oid, _ := strconv.Atoi(id)
	order := a.OrderService.GetById(oid)
	ctx.JSON(http.StatusOK, order)
}

func (a *OrderController) Query(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	//TODO 拿到当前登录用户，可把 ctx 传递到 service层
	global.Log.Infof("%#v", user)
	var orderQuery model.OrderQuery
	valid, errs := common.BindAndValid(ctx, &orderQuery)
	if !valid {
		result.FailMsg(errs.Error(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, a.OrderService.Query(orderQuery))
}
