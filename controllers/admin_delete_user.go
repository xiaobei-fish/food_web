package controllers

import "food_web/models"

type UserDeleteController struct {
	BaseController
}

func (c *UserDeleteController) Get(){
	c.Controller.EnableRender = false

	userid := c.Ctx.Input.Param(":id")
	_, _ = models.DeleteUserWithId(userid)
	models.SetUserRowsNum()

	c.Redirect("/admin/user", 302)
}
