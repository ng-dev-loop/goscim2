package models

type Userinfo struct {
	Id                 string `json:"id" xorm:"'id' not null pk unique VARCHAR(128)"`
	Password           string `json:"password" xorm:"'password' default '' VARCHAR(256)"`
	Sysadmin           int    `json:"sysadmin" xorm:"'sysadmin' default 0 INT(11)"`
	Ptzlevel           int    `json:"ptzlevel" xorm:"'ptzlevel' default 0 INT(11)"`
	ServerinfoId       string `json:"serverinfo_id" xorm:"'serverinfo_id' default '' VARCHAR(128)"`
	LogicgroupinfoId   string `json:"logicgroupinfo_id" xorm:"'logicgroupinfo_id' default '' VARCHAR(128)"`
	Maxsessionnum      int    `json:"maxsessionnum" xorm:"'maxsessionnum' default 0 INT(11)"`
	AllocateuserinfoId string `json:"allocateuserinfo_id" xorm:"'allocateuserinfo_id' default '' VARCHAR(128)"`
	Name               string `json:"name" xorm:"'name' default '' VARCHAR(128)"`
	Phones             string `json:"phones" xorm:"'phones' default '' VARCHAR(128)"`
	Email              string `json:"email" xorm:"'email' default '' VARCHAR(128)"`
	Description        string `json:"description" xorm:"'description' default '' VARCHAR(1024)"`
	Position           string `json:"position" xorm:"'position' VARCHAR(64)"`
}

type Logicgroupinfo struct {
	Id           string `json:"id" xorm:"'id' not null pk unique VARCHAR(128)"`
	Name         string `json:"name" xorm:"'name' default '' VARCHAR(128)"`
	ParentId     string `json:"parent_id" xorm:"'parent_id' default '' VARCHAR(128)"`
	Description  string `json:"description" xorm:"'description' default '' VARCHAR(1024)"`
	Abbreviation string `json:"abbreviation" xorm:"'abbreviation' VARCHAR(128)"`
}
