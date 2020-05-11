package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "myproject/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlhost := beego.AppConfig.String("mysqlhost")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqldb := beego.AppConfig.String("mysqldb")

	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册数据库直接连接
	//用户名:密码@tcp(url地址)/数据库 ,名字对应app.conf配置数据库信息
	conn := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlhost + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8"
	//注册数据库连接
	orm.RegisterDataBase("default", "mysql", conn)

	fmt.Printf("数据库连接成功！%s\n", conn)
}

func main() {
	orm.Debug = true
	beego.Run()
}

