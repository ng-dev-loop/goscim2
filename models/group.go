package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	uuid "github.com/satori/go.uuid"
	"time"
)

type SCIMGroupModel struct {
}

func (c *SCIMGroupModel) AddObject(object SCIMGroup) (objectId string) {
	c.updateObject(&object)

	if len(object.Id) == 0 {
		object.Id = uuid.Must(uuid.NewV4()).String()
	}
	created := time.Now()

	object.ExtendField.Created = created.Format(time.RFC3339)
	object.ExtendField.Schemas = append(object.ExtendField.Schemas, "urn:ietf:params:scim:schemas:core:2.0:Group")

	return object.Id
}

func (*SCIMGroupModel) GetObject(objectId string) (object *SCIMGroup, err error) {

	return nil, errors.New("ObjectId Not Exist")
}

func (*SCIMGroupModel) GetAllObject() map[string]*SCIMGroup {
	return nil
}

func (c *SCIMGroupModel) UpdateObject(objectId string, newObject *SCIMGroup) (err error) {
	c.updateObject(newObject)

	return nil
}

func (*SCIMGroupModel) DeleteObject(objectId string) {
	var user Userinfo
	_, err := DBEngine.Where("id=?", objectId).Delete(&user)
	if err != nil {
		logs.GetBeeLogger().Error("delete group db error %v", err)
		return
	}
}

func (c *SCIMGroupModel) updateObject(object *SCIMGroup) {

	var group Logicgroupinfo
	isFound, err := DBEngine.Where("id=?", object.Id).Get(&group)
	if err != nil {
		logs.GetBeeLogger().Error("AddObject db error %v", err)
		return
	}

	group.Id = object.Id
	group.Name = object.DisplayName
	group.Description = object.ExtendField.Description
	group.ParentId = object.OuUuid
	if group.ParentId == group.Id {
		group.ParentId = ""
	}

	if isFound {
		_, err = DBEngine.Where("id=?", object.Id).Update(&group)
		if err != nil {
			logs.GetBeeLogger().Error("update group fail %v", err)
		}
	} else {
		_, err := DBEngine.Insert(&group)
		if err != nil {
			logs.GetBeeLogger().Error("add group fail %v", err)
		}
	}
}
