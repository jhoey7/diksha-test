package controllers

import (
	"edot-test/repositories"
	"edot-test/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

// UserController operation
type UserController struct {
	beego.Controller
}

// Register controller for register new user.
// @Title Register controller for register new user.
// @Description Register controller for register new user.
// @Success 200 {object} models.Response
// @router /register [put]
func (c UserController) Register() {
	identifier := time.Now().UnixNano()

	userRepo := repositories.NewUserRepository(orm.NewOrm())
	svc := services.NewUserService(userRepo, identifier)
	resp := svc.Register(c.Ctx.Input.RequestBody)

	c.Data["json"] = resp
	c.ServeJSON()
	return
}

// Login controller for login user.
// @Title Login controller for login user.
// @Description Login controller for login user.
// @Success 200 {object} models.Response
// @router /login [put]
func (c UserController) Login() {
	identifier := time.Now().UnixNano()

	userRepo := repositories.NewUserRepository(orm.NewOrm())
	svc := services.NewUserService(userRepo, identifier)
	resp := svc.Login(c.Ctx.Input.RequestBody)

	c.Data["json"] = resp
	c.ServeJSON()
	return
}
