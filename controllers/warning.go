package controllers

type WarningController struct {
	BaseController
}

func (c *WarningController) Get(){
	c.TplName = "warning.html"
}
