package controllers

import (
	"fmt"
	"food_web/models"
	"strconv"
)

type SelfHomeController struct {
	BaseController
}

func (c *SelfHomeController) Get(){
	user_id := models.QueryUserWithUsername(c.Loginuser.(string))
	flag 	:= models.QueryUserMessageBool(user_id)
	genre	:= models.QueryGenre(c.Loginuser.(string))
	if flag == true{
		UserMes := models.QueryUserMessageById(strconv.Itoa(user_id))
		fmt.Println(UserMes)
		for _,u := range UserMes{
			c.Data["Words"] = u.Words
			c.Data["Picture"] = u.Picture
			c.Data["Sex"]	= u.Sex
		}
		if genre == 0 {
			c.Data["Genre"] = "买家"
		}else if genre == 1 {
			c.Data["Genre"] = "卖家"
		}else if genre == 2 {
			c.Data["Genre"] = "管理员"
		}
	}else{
		c.Data["Words"] = "未设置"
		c.Data["Sex"] 	= "未设置"
		if genre == 0 {
			c.Data["Genre"] = "买家"
		}else if genre == 1 {
			c.Data["Genre"] = "卖家"
		}else if genre == 2 {
			c.Data["Genre"] = "管理员"
		}
	}
	c.TplName = "selfhome.html"
}
