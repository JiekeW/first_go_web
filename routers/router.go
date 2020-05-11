package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/access_token/:id(\\d*)", &controllers.AccessTokenController{})
}
