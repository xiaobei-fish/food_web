package controllers

import "food_web/utils"

type BagCancelController struct {
	BaseController
}

func (c *BagCancelController) Get() {
	c.Controller.EnableRender = false

	//id为在bag表里的id
	id := c.Ctx.Input.Param(":id")

	_, _ = utils.ModifyDB("delete from user_bag where id='" + id + "'")
	c.Redirect("/bag",302)
}
