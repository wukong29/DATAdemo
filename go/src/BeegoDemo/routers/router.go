package routers

import (
	"BeegoDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/upload",&controllers.UploadController{})
}
