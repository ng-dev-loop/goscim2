package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"time"
)

type SCIMAccountModel struct {
}

func (c *SCIMAccountModel) AddObject(object SCIMAccount) string {
	c.updateObject(&object)

	created := time.Now()
	object.ExtendField.Created = created.Format(time.RFC3339)
	object.ExtendField.Schemas = append(object.ExtendField.Schemas, "urn:ietf:params:scim:schemas:core:2.0:User")

	return object.Id
}

func (*SCIMAccountModel) GetObject(objectId string) (u *SCIMAccount, err error) {
	return nil, errors.New("User not exists ")
}

func (*SCIMAccountModel) GetAllObject() map[string]*SCIMAccount {
	return nil
}

func (c *SCIMAccountModel) UpdateObject(uid string, object *SCIMAccount) (a *SCIMAccount, err error) {
	c.updateObject(object)

	return nil, nil
}

func (*SCIMAccountModel) DeleteObject(objectId string) {
	var user Userinfo
	_, err := DBEngine.Where("id=?", objectId).Delete(&user)
	if err != nil {
		logs.GetBeeLogger().Error("AddObject db error %v", err)
		return
	}
}

func (c *SCIMAccountModel) updateObject(object *SCIMAccount) {

	var user Userinfo
	isFound, err := DBEngine.Where("id=?", object.UserName).Get(&user)
	if err != nil {
		logs.GetBeeLogger().Error("AddObject db error %v", err)
		return
	}

	if object.UserName == "admin" {
		logs.GetBeeLogger().Warn("user admin not sync")
		return
	}

	user.Id = object.UserName
	user.Name = object.DisplayName
	user.Description = object.Id
	user.Sysadmin = 16
	user.Ptzlevel = 1
	user.AllocateuserinfoId = "admin"
	if len(object.Emails) > 0 {
		user.Email = object.Emails[0].Value
	}
	if len(object.PhoneNumbers) > 0 {
		user.Phones = object.PhoneNumbers[0].Value
	}
	if len(object.Belongs) > 0 {
		/*isContains := false
		for _, obj := range object.Belongs {
			if obj.BelongOuUuid == user.LogicgroupinfoId {
				isContains = true
				break
			}
		}
		if isContains == false {
			user.LogicgroupinfoId = object.Belongs[0].BelongOuUuid
		}*/
		user.LogicgroupinfoId = object.Belongs[0].BelongOuUuid
	}

	if isFound {
		_, err = DBEngine.Where("id=?", user.Id).Update(&user)
		if err != nil {
			logs.GetBeeLogger().Error("update user fail %v", err)
		}
	} else {
		_, err := DBEngine.Insert(&user)
		if err != nil {
			logs.GetBeeLogger().Error("add user fail %v", err)
		}
	}
}
