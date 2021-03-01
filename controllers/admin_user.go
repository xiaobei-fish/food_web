package controllers

import (
	"fmt"
	"food_web/models"
	"strconv"
)

type AdminUserController struct {
	BaseController
}

func (c *AdminUserController) Get(){
	page, _ := c.GetInt("page")
	var uList []models.User

	if page <= 0 {
		page = 1
	}
	//设置分页
	uList, _ = models.FindUserWithPage(page)

	c.Data["PageCode"] = models.ConfigAdminFooterPageCode(page)
	c.Data["HasFooter"] = true

	fmt.Println("IsLogin:", c.IsLogin, c.Loginuser)
	c.Data["Content"] = models.MakeAdminUserBlocks(uList, c.IsLogin)

	c.TplName = "admin_user.html"
}

func (c *AdminUserController) Post(){
	money := c.GetString("money")
	fmt.Println(money)
	var MoneyList []models.UserMoney
	var count float64
	var lmoney float64
	MoneyList, _ = models.QueryUserMoney(c.Loginuser.(string))

	for _,mon := range MoneyList{
		lmoney, _ = strconv.ParseFloat(mon.Money,64)
		count += lmoney
	}

	tmp, _ := strconv.ParseFloat(money,64)
	cnt := strconv.FormatFloat(tmp + count,'E',-1,64)
	_, _ = models.UpdateUserMoney(c.Loginuser.(string), cnt)

	c.Data["json"] = map[string]interface{}{"code": 1, "message": "充值成功"}

	c.ServeJSON()
}
