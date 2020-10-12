package main

import (
	"BeegoProject0603/db_mysql"
	_ "BeegoProject0603/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.Connect()
	beego.Run()
}

