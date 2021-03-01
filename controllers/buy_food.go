package controllers

import (
	"fmt"
	"food_web/models"
	"strconv"
)

type BuyFoodController struct {
	BaseController
}

func (c *BuyFoodController) Get(){
	c.Controller.EnableRender = false

	var moneyList []models.UserMoney
	moneyList, _ = models.QueryUserMoney(c.Loginuser.(string))
	var lmoney float64
	for _,money := range moneyList{
		lmoney,_ = strconv.ParseFloat(money.Money,64)
	}
	var foodList []models.Food_info
	var cost float64
	foodid := c.Ctx.Input.Param(":id")
	foodList = models.QueryFoodById(foodid)
	for _,food := range foodList{
		cost,_ = strconv.ParseFloat(food.FoodPrice,64)
	}

	if lmoney >= cost {
		//删除上架货物，加入到该用户背包，扣款
		for _,food := range foodList{
			userid := models.QueryUserWithUsername(c.Loginuser.(string))
			_, _ = models.AddFoodToBag(food,userid)
			_, _ = models.DeleteFoodWithId(strconv.Itoa(food.Id))
			money := strconv.FormatFloat(lmoney-cost,'E',-1,64)
			_, _ = models.UpdateUserMoney(c.Loginuser.(string), money)
		}
		fmt.Println("购买成功")
		c.Redirect("/shop_car", 302)
	}else{
		fmt.Println("余额不足")
		c.Redirect("/warning", 302)
	}
}
