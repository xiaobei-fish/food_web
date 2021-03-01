package controllers

import (
	"food_web/models"
	"strconv"
)

type UserBagController struct {
	BaseController
}

func (c *UserBagController) Get(){
	username := c.Loginuser.(string)
	userid 	 := models.QueryUserWithUsername(username)

	var bagList []models.UserBag
	bagList, _ = models.QueryUserBag(strconv.Itoa(userid))

	c.Data["Content"] = models.MakeBagBlocks(bagList, c.IsLogin)

	c.TplName = "bag.html"
}
