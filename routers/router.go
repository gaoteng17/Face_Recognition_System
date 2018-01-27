package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    //beego.Router("/", &controllers.IdentityController{})
    beego.Router("/img", &controllers.ImgController{})
}
