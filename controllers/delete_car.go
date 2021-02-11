package controllers

import (
	"food_web/models"
	"log"
	"strconv"
)

type DeleteCarController struct {
	BaseController
}

func (c *DeleteCarController) Get(){
	c.Controller.EnableRender = false
	foodid := c.Ctx.Input.Param(":id")
	userid := models.QueryUserWithUsername(c.Loginuser.(string))
	userId := strconv.Itoa(userid)

	_, err := models.DeleteCarFoodWithId(foodid,userId)
	if err != nil{
		log.Println(err)
		panic(err)
	}


	c.Redirect("/shop_car", 302)
}
