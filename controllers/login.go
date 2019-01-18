package controllers

import (
	"fmt"
	"gameadmin/models"
	"gameadmin/utils"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Login() {
	name := this.GetString("username")
	pwd  := this.GetString("password")
	if name == "" || pwd == "" {
		this.Ctx.WriteString("参数有误")
		return
	}
	user, err := models.GetOneUser(fmt.Sprintf("name=%s",name))
	if err != nil {
		this.Ctx.WriteString("该账号还未注册")
		return
	}

	if user.Password != utils.Stringmd5(pwd) {
		this.Ctx.WriteString("密码不正确")
		return
	}


	this.Data["user"]= name
	users,err  := models.GetAllUsers()

	this.Data["users"]=users
	if len(users) > defaultPageSize {
		this.Data["users"]=users[0:defaultPageSize]
	}
	this.Data["page"]= utils.PageUtil(len(users),1,defaultPageSize)


	this.LayoutSections = make(map[string]string,0)
	this.LayoutSections["headercss"]="layout/header_css.html"
	this.LayoutSections["footerjs"]="layout/footer_js.html"
	this.Layout = "layout/layout_base.html"
	this.TplName ="user_table.html"
	return
}



func (this *LoginController) Logout()  {

}