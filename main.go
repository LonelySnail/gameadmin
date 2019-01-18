package main

import (
	_ "gameadmin/routers"
	_ "gameadmin/utils"
	"github.com/astaxie/beego"
)

func main() {
	//TODO  验证过滤路由 判断玩家是否登陆
	beego.Run()
}

