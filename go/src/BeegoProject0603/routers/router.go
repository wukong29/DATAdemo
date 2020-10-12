package routers

import (
	"BeegoProject0603/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register",&controllers.RegisterController{})
	//http://127.0.0.1:8080
	beego.Router("/", &controllers.MainController{})
}
