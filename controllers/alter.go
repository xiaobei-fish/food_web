package controllers

import (
	"fmt"
	"food_web/models"
	"github.com/astaxie/goredis"
	"strings"
)

type AlterController struct {
	BaseController
}

func (c *AlterController) Get(){
	c.TplName = "alter.html"
}

func (c *AlterController) Post(){
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	//获取浏览器传输的数据，通过表单的name属性获取值
	sex 	:= c.GetString("sex")
	words	:= c.GetString("words")

	Pic, err := client.Get("Head")
	if err != nil {
		panic(err)
	}
	pic := fmt.Sprintf("%s",Pic)
	pic = strings.Replace(pic,"\\","/",-1)

	id := models.QueryUserWithUsername(c.Loginuser.(string))
	//存入数据库中
	mes := models.UserMessage{Words:words, Picture:"../" + pic, Sex:sex, UserId:id}
	_, _ = models.AddMes(mes)

	c.Data["json"] = map[string]interface{}{"code": 1, "message":"上传成功"}
	c.ServeJSON()
	c.Redirect("/selfhome", 302)
}
