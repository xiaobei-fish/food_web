package controllers

import(
	"fmt"
	"food_web/models"
	_ "github.com/astaxie/beego"
)

type HomeController struct {
	//beego.Controller
	BaseController
}

func(c *HomeController) Get(){
	page, _ := c.GetInt("page")
	var foodList []models.Food_info

	if page <= 0 {
		page = 1
	}
	//设置分页
	foodList, _ = models.FindFoodWithPage(page)

	c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	c.Data["HasFooter"] = true

	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.Data["Content"] = models.MakeHomeBlocks(foodList, c.IsLogin)
	c.TplName = "home.html"
}
