package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/ng-dev/goscim2/models"
	"net/http"

	"github.com/astaxie/beego"
)

// Operations about Users
type AccountController struct {
	beego.Controller
	accountModel models.SCIMAccountModel
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.SCIMAccount	true	"body for user content"
// @Success 200 {string} models.SCIMAccount.Id
// @Failure 403 body is empty
// @router / [post]
func (u *AccountController) Post() {
	var user models.SCIMAccount
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.GetBeeLogger().Error("add account %v", err)
		u.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	uid := u.accountModel.AddObject(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.SCIMAccount
// @router / [get]
func (u *AccountController) GetAll() {
	users := u.accountModel.GetAllObject()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SCIMAccount
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *AccountController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := u.accountModel.GetObject(uid)
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
func (u *AccountController) Put() {
	var user models.SCIMAccount
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.GetBeeLogger().Error("update account %v", err)
		u.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(user.Id) == 0 {
		logs.GetBeeLogger().Error("update account no user id")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = u.accountModel.UpdateObject(user.Id, &user)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = JsonHttpResult{
			ErrorNumber: 0,
			Errors: []string{
				fmt.Sprintf("update success %v", user.UserName),
			},
		}
	}

	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router / [delete]
func (u *AccountController) Delete() {
	// ?id=6498786832594607257&applicationUsername=ceshirenyuan3&ddAccountId=&udUsername=ceshirenyuan3
	uid := u.GetString("id")
	udUsername := u.GetString("udUsername")
	logs.GetBeeLogger().Debug("delete %v %v %v", uid, udUsername, u.Ctx.Request.RequestURI)

	u.accountModel.DeleteObject(udUsername)

	result := JsonHttpResult{
		ErrorNumber: 0,
		Errors: []string{
			"delete success",
		},
	}

	u.Data["json"] = result
	u.ServeJSON()
}
