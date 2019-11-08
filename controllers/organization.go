package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ng-dev/goscim2/models"
	"net/http"
)

// Operations about Users
type OrganizationController struct {
	beego.Controller
	organizationModel models.SCIMOrganizationModel
}

// @Title create organization
// @Description create organization
// @Param	body		body 	models.SCIMOrganization	true		"body for user content"
// @Success 200 {string} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *OrganizationController) Post() {
	var object models.SCIMOrganization
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &object)
	if err != nil {
		logs.GetBeeLogger().Error("add organization %v", err)
		u.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	uid := u.organizationModel.AddObject(object)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.SCIMAccount
// @router / [get]
func (u *OrganizationController) GetAll() {
	objects := u.organizationModel.GetAllObject()
	u.Data["json"] = objects
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SCIMAccount
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *OrganizationController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := u.organizationModel.GetObject(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.SCIMAccount	true		"body for user content"
// @Success 200 {object} models.SCIMAccount
// @Failure 403 :uid is not int
// @router / [put]
func (u *OrganizationController) Put() {
	result := JsonHttpResult{
		ErrorNumber: 400,
		Errors: []string{
			"update success",
		},
	}

	var object models.SCIMOrganization
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &object)
	if err != nil {
		result.ErrorNumber = 500
		result.Errors = []string{fmt.Sprintf("update fail json un marshal %v", err)}
	} else {
		err = u.organizationModel.UpdateObject(object.OrganizationUuid, &object)
		if err != nil {
			result.Errors = []string{err.Error()}
		} else {
			result.ErrorNumber = 0
			result.Errors = []string{"update success"}
		}
	}

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router / [delete]
func (u *OrganizationController) Delete() {
	uid := u.GetString("id")
	u.organizationModel.DeleteObject(uid)

	result := JsonHttpResult{
		ErrorNumber: 0,
		Errors: []string{
			"delete success",
		},
	}

	u.Data["json"] = result
	u.ServeJSON()
}
