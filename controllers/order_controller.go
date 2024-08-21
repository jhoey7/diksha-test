package controllers

import (
	"edot-test/repositories"
	"edot-test/services"
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
	productDtlRepo := repositories.NewProductDetailRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	orderDtlRepo := repositories.NewOrderDetailRepository(db)
	svc := services.NewCheckoutOrderService(productRepo, productDtlRepo, orderRepo, orderDtlRepo, db, identifier)
	resp := svc.Checkout(o.Ctx.Input.RequestBody)

	o.Data["json"] = resp
	o.ServeJSON()
	return
}

func (o OrderController) Confirm() {
	identifier := time.Now().UnixNano()

	db := orm.NewOrm()
	productRepo := repositories.NewProductRepository(db)
	productDtlRepo := repositories.NewProductDetailRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	orderDtlRepo := repositories.NewOrderDetailRepository(db)
	svc := services.NewCheckoutOrderService(productRepo, productDtlRepo, orderRepo, orderDtlRepo, db, identifier)
	resp := svc.Confirm(o.Ctx.Input.Param(":pubId"))

	o.Data["json"] = resp
	o.ServeJSON()
	return
}
