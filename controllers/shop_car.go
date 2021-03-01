package controllers

import (
	"fmt"
	"food_web/models"
)

type ShopcarController struct {
	BaseController
}

func (c *ShopcarController) Get(){
	var foodList []models.Food_info
	var cost float64

	foodList, _, cost = models.FindCarWithUserId(c.Loginuser.(string))
	//fmt.Println(foodList)

	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.Data["Content"] = models.MakeCarBlocks(foodList, c.IsLogin)
	c.Data["Cost"]    = cost

	c.TplName = "shop_car.html"
}
