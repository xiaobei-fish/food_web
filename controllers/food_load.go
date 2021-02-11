package controllers

import (
	"fmt"
	"food_web/models"
	"github.com/astaxie/goredis"
	"strconv"
	"strings"
)

type FoodLoadController struct {
	BaseController
}

func (c *FoodLoadController) Get(){
	c.TplName = "food_load.html"
}

func (c *FoodLoadController) Post() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	//获取浏览器传输的数据，通过表单的name属性获取值
	store 	:= c.GetString("store")
	price 	:= c.GetString("price")
	intro 	:= c.GetString("intro")

	Pic, err := client.Get("Pic")
	if err != nil {
		panic(err)
	}
	pic := fmt.Sprintf("%s",Pic)
	pic = strings.Replace(pic,"\\","/",-1)

	//存入数据库中
	food := models.Food_info{FoodStore: store, FoodPic:"../" + pic, FoodIntro: intro, FoodPrice: price}
	_, _ = models.AddFood(&food)
	models.SetFoodRowsNum()
	constr := fmt.Sprintf("where food_store='%s'",store)
	foodid := models.QueryFoodWithStore(constr)
	var seller models.Seller
	seller.Username = c.Loginuser.(string)
	seller.FoodId 	= strconv.Itoa(foodid)
	_, _ = models.InsertSeller(seller)

	c.Data["json"] = map[string]interface{}{"code": 1, "message":"上传成功"}
	c.ServeJSON()
	c.Redirect("/food/add", 302)
}
