package routers

import (
	"gameadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Index")
    beego.Router("/register",&controllers.MainController{},"*:Register")
	beego.Router("/doRegister",&controllers.MainController{},"*:DoRegister")
    beego.Router("/login",&controllers.LoginController{},"*:Login")
	beego.Router("/getUserInfo",&controllers.UserController{},"GET:GetAllUsersInfo")
	beego.Router("/user",&controllers.UserController{},"GET:GetUsersInfo")
	beego.Router("/test",&controllers.UserController{},"GET:Test")
}
