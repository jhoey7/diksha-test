package controllers

import (
	"diskha-test/models"
	"diskha-test/repositories"
	"diskha-test/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

// ProductController operation
type ProductController struct {
	beego.Controller
}

// List controller for get product list.
// @Title List controller for get product list.
// @Description List controller for get product list.
// @Success 200 {object} models.Response
// @router [get]
func (p ProductController) List() {
	identifier := time.Now().UnixNano()

	var req models.ProductListRequest

	req.Keyword = p.GetString("keyword", "")
	req.Order = p.GetString("order", "createdAt")
	req.Sort = p.GetString("sort", "DESC")
	req.Page, _ = p.GetInt("page", 0)
	req.Size, _ = p.GetInt("size", 20)

	productRepo := repositories.NewProductRepository(orm.NewOrm())
	svc := services.NewListProductService(productRepo, identifier)
	resp := svc.List(req)

	p.Data["json"] = resp
	p.ServeJSON()
	return
}
