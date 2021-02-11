package controllers

import (
	"fmt"
	"food_web/models"
	"github.com/astaxie/goredis"
	"strconv"
)

type CollectController struct {
	DetailController
	BaseController
}

func (c *CollectController) Get(){
	c.Controller.EnableRender = false
	//得到要收藏食品的信息
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	userName 	:= c.Loginuser
	foodId, Err	:=	client.Get("foodid")
	if Err != nil {
		panic(Err)
	}

	fmt.Println("user:",userName)
	fmt.Println("food:",foodId)

	//检查是否已经在收藏夹中
	userid := models.QueryUserWithUsername(userName.(string))
	id := models.QueryFoodWithUserId(strconv.Itoa(userid), string(foodId))
	fmt.Println("id:", id)
	if id > 0 {
		fmt.Println("添加失败")
		c.Redirect("/food/" + string(foodId), 302)
		return
	}
	//实例化model，将它存入到数据库中
	model 	:= models.Collect{Userid:strconv.Itoa(userid),Foodid:string(foodId)}
	_, err 	:= models.CollectFood(model)
	//返回数据给浏览器
	if err == nil {
		//无误
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}
	c.Redirect("/food/" + string(foodId), 302)
}