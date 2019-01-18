package utils


import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/astaxie/beego"
	"github.com/weikaishio/go-logger/logger"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlDriver map[string]*gorm.DB
func init()   {
	dbType := beego.AppConfig.String("db_type")
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	dbName := beego.AppConfig.String(dbType + "::db_name")
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	// dbCharset := beego.AppConfig.String(dbType + "::db_charset")
	MysqlDriver = make(map[string]*gorm.DB)
	db, err := gorm.Open("mysql", dbUser+":"+dbPwd+"@("+dbHost+")/"+dbName+"?charset=utf8&parseTime=False&loc=Local")
	if err != nil {
		db.DB().Close()
		log.Fatal(fmt.Sprintf("NewMysqlDriver: open Database %s error %v:", dbHost, err))
		return
	}

	db.LogMode(true)
	MysqlDriver[dbAlias]=db
	return
}

func GetDefaultDb()  *gorm.DB{
	dbType := beego.AppConfig.String("db_type")
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	db, ok := MysqlDriver[dbAlias]
	if !ok {
		logger.LogWarn("GetDefaultDb: not register db:",dbAlias)
		return nil
	}
	return db
}

func GetMysqlDb(name string) *gorm.DB {
	db, ok := MysqlDriver[name]
	if !ok {
		return nil
	}
	return db
}