package controllers

import (
	"food_web/models"
)

type AdminFoodController struct {
	BaseController
}

func (c *AdminFoodController) Get(){
	page, _ := c.GetInt("page")
	var foodList []models.Food_info

	if page <= 0 {
		page = 1
	}
	//设置分页
	foodList, _ = models.FindFoodWithPage(page)

	c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	c.Data["HasFooter"] = true

	c.Data["Content"] = models.MakeAdminFoodBlocks(foodList, c.IsLogin)

	c.TplName = "admin_food.html"
}
