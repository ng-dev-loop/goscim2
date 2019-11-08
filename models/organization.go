package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"time"
)

type SCIMOrganizationModel struct {
}

func (c *SCIMOrganizationModel) AddObject(object SCIMOrganization) (objectId string) {
	c.updateObject(&object)

	if len(object.OrganizationUuid) == 0 {
		logs.GetBeeLogger().Error("error organization OrganizationUuid")
		return ""
	}
	created := time.Now()

	object.ExtendField.Created = created.Format(time.RFC3339)
	object.ExtendField.Schemas = append(object.ExtendField.Schemas, "urn:ietf:params:scim:schemas:core:2.0:Organization")

	return object.OrganizationUuid
}

func (*SCIMOrganizationModel) GetObject(objectId string) (object *SCIMOrganization, err error) {

	return nil, errors.New("ObjectId Not Exist")
}

func (*SCIMOrganizationModel) GetAllObject() map[string]*SCIMOrganization {
	return nil
}

func (c *SCIMOrganizationModel) UpdateObject(objectId string, object *SCIMOrganization) (err error) {
	c.updateObject(object)

	return nil
}

func (*SCIMOrganizationModel) DeleteObject(objectId string) {
	help := SCIMGroupModel{}
	help.DeleteObject(objectId)
}

func (*SCIMOrganizationModel) updateObject(object *SCIMOrganization) {
	group := SCIMGroup{
		Id:          object.OrganizationUuid,
		DisplayName: object.Organization,
		OuUuid:      object.ParentUuid,
		ExtendField: SCIMExtendField{
			Description: object.Description,
		},
	}

	help := SCIMGroupModel{}
	help.updateObject(&group)
}
