package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["data"] = "ty2xy"
	c.TplName = "test.html"
}
func (c *MainController) Post() {
	c.Data["data"] = "ty2xy1314"
	c.TplName = "test.html"
}
