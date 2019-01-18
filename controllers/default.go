package controllers

import (
	"fmt"
	"gameadmin/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	app := beego.AppConfig.String("site.app")
	name := beego.AppConfig.String("site.name")
	this.Data["title"] =app  + name
	this.Data["siteApp"] = app
	this.Data["siteName"]= name
	this.Data["siteVersion"] = beego.AppConfig.String("site.version")
	this.TplName = "index.html"
}


func (this *MainController)Register()  {
	this.TplName= "register.html"
}

func (this *MainController)DoRegister()  {
	name := this.GetString("name")
	pwd := this.GetString("password")
	rePwd := this.GetString("rePassword")
	phone := this.GetString("phone")
	if name == "" || pwd == "" || rePwd == "" || phone=="" {
		this.Ctx.WriteString("参数有误")
		return
	}

	if pwd != rePwd {
		this.Ctx.WriteString("两次密码不一致，请重新shur")
		return
	}

	user,_ := models.GetOneUser(fmt.Sprintf("name=%s",name))
	if user.Name == name {
		this.Ctx.WriteString("该账号已经注册")
		return
	}

	err := models.NewUser(name,pwd,phone)
	if err != nil {
		this.Ctx.WriteString("注册失败")
		return
	}
	this.Ctx.WriteString("success")
	return
}