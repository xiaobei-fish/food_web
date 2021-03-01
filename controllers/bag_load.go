package controllers

import (
	"food_web/models"
	"food_web/utils"
)

type BagLoadController struct {
	BaseController
}

func (c *BagLoadController) Get(){
	c.Controller.EnableRender = false

	//id为在bag表里的id
	id := c.Ctx.Input.Param(":id")

	var bagList []models.UserBag
	bagList,_ = models.QueryUserBagById(id)

	for _,bag := range bagList {
		_, _ = utils.ModifyDB("insert into food_info(food_store,food_pic,food_intro,food_price) values(?,?,?,?)",
		bag.FoodStore,bag.FoodPic,bag.FoodIntro,bag.FoodPrice)
		_, _ = utils.ModifyDB("delete from user_bag where id='" + id + "'")
	}
	c.Redirect("/bag",302)
	//前面没说清楚，现在要做这些，上架之后还要捆绑用户id，然后虚拟货币更新
	//数据库重设麻烦，这个功能不写了
}
