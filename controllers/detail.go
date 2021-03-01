package controllers

import (
	"fmt"
	"food_web/models"
	"food_web/utils"
	"github.com/astaxie/goredis"
	"log"
)

type DetailController struct {
	BaseController
}

func (c *DetailController) Get(){
	var foodlist []models.Food_info
	foodid := c.Ctx.Input.Param(":id")
	fmt.Printf("foodid:%s\n",foodid)
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("foodid",[]byte(foodid))
	if err != nil {
		log.Println(err)
	}

	foodlist = models.QueryFoodById(foodid)
	//fmt.Println(foodlist)
	for _,food := range foodlist{
		c.Data["Id"]	=foodid
		c.Data["Intro"] = food.FoodIntro
		c.Data["Pic"] = food.FoodPic
		c.Data["Price"] = food.FoodPrice
		c.Data["Store"] = food.FoodStore
	}
	var commentList []models.Comment
	commentList = models.QueryFoodCommentById(foodid)
	fmt.Println(commentList)

	c.Data["Content"] = models.MakeCommentBlocks(commentList,c.IsLogin)

	c.TplName = "food_detail.html"
}

func (c *DetailController) Post(){
	comment := c.GetString("comment")

	foodid := c.GetString("id")

	_, _ = utils.ModifyDB("insert into comment(username,foodid,intro) values(?,?,?)",
		c.Loginuser.(string),foodid,comment)

	c.Data["json"] = map[string]interface{}{"code": 1, "message": "评论成功"}
	c.ServeJSON()
}
