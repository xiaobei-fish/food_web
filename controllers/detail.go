package controllers

import (
	"fmt"
	"food_web/models"
	"github.com/astaxie/goredis"
	"log"
)

type DetailController struct {
	BaseController
}

func (c *DetailController) Get(){
	var foodlist []models.Food_info
	foodid := c.Ctx.Input.Param(":id")
	fmt.Printf("foodid:%s\n",foodid)
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("foodid",[]byte(foodid))
	if err != nil {
		log.Println(err)
	}

	foodlist = models.QueryFoodById(foodid)
	//fmt.Println(foodlist)
	for _,food := range foodlist{
		c.Data["Id"]	=foodid
		c.Data["Intro"] = food.FoodIntro
		c.Data["Pic"] = food.FoodPic
		c.Data["Price"] = food.FoodPrice
		c.Data["Store"] = food.FoodStore
	}
	c.TplName = "food_detail.html"
}
