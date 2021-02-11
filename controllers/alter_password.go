package controllers

import (
	"fmt"
	"food_web/models"
)

type AltPasswordController struct {
	BaseController
}

func (c *AltPasswordController) Get(){
	c.TplName = "alter_password.html"
}

func (c *AltPasswordController) Post(){
	old  := c.GetString("oldpassword")
	news := c.GetString("newpassword")

	flag := models.TestOldPassword(old)
	if flag > 0 {
		_, _ = models.AlterPassword(c.Loginuser.(string), news)
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "修改成功,请重新登录"}
		c.DelSession("loginuser")
		fmt.Println("修改成功")
		
		c.ServeJSON()
		c.Redirect("/", 302)
		return
	}else{
		if old == news {
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改失败,旧密码输入错误"}
			fmt.Println("修改失败,旧密码输入错误")

			c.ServeJSON()
			c.Redirect("/", 302)
		}else{
			c.Data["json"] = map[string]interface{}{"code": 0, "message": "修改失败,新旧密码一致"}
			fmt.Println("修改失败,新旧密码一致")

			c.ServeJSON()
			c.Redirect("/", 302)
		}
	}
}
