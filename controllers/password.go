package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ng-dev-loop/goscim2/models"
)

// Operations about Users
type PasswordController struct {
	beego.Controller
	organizationModel models.SCIMOrganizationModel
}

// @Title create organization
// @Description create organization
// @Param	body		body 	models.SCIMOrganization	true		"body for user content"
// @Success 200 {string} models.SCIMAccount.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PasswordController) Post() {

	fmt.Println(string(u.Ctx.Input.RequestBody))

	u.Data["json"] = ""
	u.ServeJSON()
}
