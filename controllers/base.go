package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"food_web/models"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	IsSeller  bool
	Loginuser interface{}
}

//判断是否登录
/*
	这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，
    用户可以重写这个函数实现类似用户验证之类。
*/
func (c *BaseController) Prepare() {
	loginuser := c.GetSession("loginuser")
	fmt.Println("loginuser---->", loginuser)
	username := fmt.Sprintf("%s",loginuser)
	if loginuser != nil {
		c.IsLogin = true
		c.Loginuser = loginuser
		c.IsSeller  = models.QueryUserGenre(username)
	} else {
		c.IsLogin  = false
		c.IsSeller = false
	}
	c.Data["Username"] = loginuser
	c.Data["IsLogin"]  = c.IsLogin
	c.Data["IsSeller"] = c.IsSeller
}
