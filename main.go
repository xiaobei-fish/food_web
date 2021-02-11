package main

import (
	_ "food_web/routers"
	"food_web/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	utils.LogOn()
	beego.Run()
}

