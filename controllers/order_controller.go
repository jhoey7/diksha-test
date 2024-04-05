package controllers

import (
	"diskha-test/repositories"
	"diskha-test/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

// OrderController operation
type OrderController struct {
	beego.Controller
}

// Checkout controller for checkout order.
// @Title Checkout controller for checkout order.
// @Description Checkout controller for checkout order.
// @Success 200 {object} models.Response
// @router /checkout [get]
func (o OrderController) Checkout() {
	identifier := time.Now().UnixNano()

	db := orm.NewOrm()
	productRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	orderDtlRepo := repositories.NewOrderDetailRepository(db)
	svc := services.NewCheckoutOrderService(productRepo, orderRepo, orderDtlRepo, db, identifier)
	resp := svc.Checkout(o.Ctx.Input.RequestBody)

	o.Data["json"] = resp
	o.ServeJSON()
	return
}
