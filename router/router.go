package routers

import (
	"edot-test/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/1.0",
			beego.NSNamespace("/users",
				beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
				beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			),

			beego.NSNamespace("/products",
				beego.NSRouter("", &controllers.ProductController{}, "get:List"),
			),

			beego.NSNamespace("/orders",
				beego.NSRouter("/checkout", &controllers.OrderController{}, "post:Checkout"),
				beego.NSRouter("/:pubId/confirm", &controllers.OrderController{}, "patch:Confirm"),
			),
		)

	beego.AddNamespace(ns)
}
