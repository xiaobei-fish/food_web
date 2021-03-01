package controllers

import "food_web/models"

type FoodDeleteController struct {
	BaseController
}

func (c *FoodDeleteController) Get(){
	c.Controller.EnableRender = false

	foodid := c.Ctx.Input.Param(":id")
	_, _ = models.DeleteFoodWithId(foodid)
	models.SetFoodRowsNum()

	c.Redirect("/admin/food", 302)
}
