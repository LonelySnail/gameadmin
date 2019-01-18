package models

import (
	"gameadmin/utils"
	"github.com/astaxie/beego/orm"
	"time"
)

type Users struct{
	Id  		int 	`orm:"auto" json:"id"`
	Name 		string	`orm:"size(255)" json:"name"`
	Password	string 	`orm:"size(255)" json:"password"`
	Phone		string		`orm:"size(15)" json:"phone"`
	Regtime     int64 	`orm:"size(15)" json:"regtime"` //注册时间
	LastTime    int64	`json:"last_time"`	//最后一次登录时间
}

func init() {
	orm.RegisterModel(new(Users))
}

func NewUser(name, pwd,phone  string)  error{
	now := time.Now().Unix()
	// o := orm.NewOrm()
	user := new(Users)
	user.Name = name
	user.Password = utils.Stringmd5(pwd)
	user.Phone = phone
	user.Regtime = now
	user.LastTime =now
	db := utils.GetDefaultDb()
	res := db.Create(user)

	// _,err := o.Insert(user)
	return res.Error
}

func GetOneUser(sql interface{})  (*Users,error){
	db := utils.GetDefaultDb()
	user := new(Users)
	res := db.Where(sql).Last(user)
	return user,res.Error
}

func GetAllUsers()([]Users,error)  {
	var users []Users
	db := utils.GetDefaultDb()
	res := db.Find(&users)
	return users,res.Error
}

func GetUsersBySql(sql interface{})([]Users,error)  {
	var users []Users
	db := utils.GetDefaultDb()
	res := db.Where(sql).Find(&users)
	return users,res.Error
}

func GetUsersByLimit(limit,offset interface{}) ([]Users,error) {
	var users []Users
	db := utils.GetDefaultDb()

	res := db.Offset(offset).Limit(limit).Find(&users)
	return users,res.Error
}

func UpdateUser(name interface{})  {

}