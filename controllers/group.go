package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ng-dev/goscim2/models"
	"net/http"
)

// Operations about group
type GroupController struct {
	beego.Controller
	groupModel models.SCIMGroupModel
}

// @Title Create
// @Description create group
// @Param	body		body 	models.SCIMGroup	true		"The object content"
// @Success 200 {string} models.SCIMGroup.Id
// @Failure 403 body is empty
// @router / [post]
func (o *GroupController) Post() {

	var ob models.SCIMGroup
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		logs.GetBeeLogger().Error("add group, unmarshal fail %v", err)
		o.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	groupId := o.groupModel.AddObject(ob)
	o.Data["json"] = map[string]string{"ObjectId": groupId}
	o.ServeJSON()
}

// @Title Get
// @Description find object by groupId
// @Param	objectId		path 	string	true		"the groupId you want to get"
// @Success 200 {object} models.SCIMGroup
// @Failure 403 :groupId is empty
// @router /:groupId [get]
func (o *GroupController) Get() {
	groupId := o.Ctx.Input.Param(":groupId")
	if groupId != "" {
		ob, err := o.groupModel.GetObject(groupId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.SCIMGroup
// @Failure 403 :groupId is empty
// @router / [get]
func (o *GroupController) GetAll() {
	obs := o.groupModel.GetAllObject()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The groupId you want to update"
// @Param	body		body 	models.SCIMGroup	true		"The body"
// @Success 200 {object} models.SCIMGroup
// @Failure 403 :groupId is empty
// @router / [put]
func (o *GroupController) Put() {
	result := JsonHttpResult{
		ErrorNumber: 400,
		Errors: []string{
			"delete fail",
		},
	}

	var ob models.SCIMGroup
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	if err != nil {
		result.ErrorNumber = 500
		result.Errors = []string{fmt.Sprintf("update fail json un marshal %v", err)}
	} else {
		err := o.groupModel.UpdateObject(ob.Id, &ob)
		if err != nil {
			result.ErrorNumber = 500
			result.Errors = []string{fmt.Sprintf("update fail %v", err)}
		} else {
			result.ErrorNumber = 0
			result.Errors = []string{"delete success"}
		}
	}

	o.Data["json"] = result
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 groupId is empty
// @router / [delete]
func (o *GroupController) Delete() {
	objectId := o.Ctx.Input.Param(":groupId")
	o.groupModel.DeleteObject(objectId)

	result := JsonHttpResult{
		ErrorNumber: 0,
		Errors: []string{
			"delete success",
		},
	}

	o.Data["json"] = result
	o.ServeJSON()
}
