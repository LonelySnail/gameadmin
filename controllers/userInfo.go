package controllers

import (
	"fmt"
	"gameadmin/models"
	"gameadmin/utils"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func(this *UserController)Test ()  {
	fmt.Println("****")
	users, _ := models.GetAllUsers()
	this.Data["title"]="222"
	this.Data["total"]=14
	this.Data["rows"]=users
	this.TplName ="test-table.html"


}


func(this *UserController) GetAllUsersInfo ()   {
	users, _ := models.GetAllUsers()

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
}

func (this *UserController) GetUsersInfo()  {
	page, _:= this.GetInt("page")
	pageSize,_ := this.GetInt("pageSize")
	users,_ := models.GetUsersByLimit(pageSize,(page-1)*pageSize)
	this.Data["users"]=users
	this.TplName ="user_table.html"
	return
}

