package models

// https://hytera.idsmanager.com/enduser/developer/docs/ud#doc_con_05_01_05

// 组织结构
type SCIMOrganization struct {
	Description      string          `json:"description"`
	Organization     string          `json:"organization"`     // >=1 组织机构名称
	ParentUuid       string          `json:"parentUuid"`       // >=1 所属的父级组织机构的唯一id,该id是SP同步过来的，所以在IPG中称为父级外部id。必填
	Manager          []SCIMManager   `json:"manager"`          // 管理者,可为多个
	OrganizationUuid string          `json:"organizationUuid"` // ou外部id(唯一)  组织机构的唯一id,该id是SP同步过来的，所以在IPG中称为外部id。必填
	RegionId         string          `json:"regionId"`         // 区域id 选填
	RootNode         bool            `json:"rootNode"`         // 是否rootNode 选填
	Type             string          `json:"type"`             /* 组织机构或部门 type为SELF_OU(自建组织机构)时有可能会有值,可为空； type为DEPARTMENT("自建部门")不会出现值*/
	LevelNumber      int             `json:"levelNumber"`      // int  机构排序号
	Status           int             `json:"status"`
	ChildrenOuUuid   []string        `json:"childrenOuUuid"` // OU的所有直属子集  有必填，为了保证OU的结构
	ExtendField      SCIMExtendField `json:"extendField"`    // 用于存放扩展字段的对象	如有扩展字段，则该属性必填
}

// 组
type SCIMGroup struct {
	Id          string          `json:"id"`          // 账户组id,唯一,可为空
	DisplayName string          `json:"displayName"` // 组显示名称
	OuUuid      string          `json:"ouUuid"`      // 	>=43	所属组织单位(OU)的外部ID
	Members     []SCIMMembers   `json:"members"`     // 组成员,已经存在的账户外部ID和账户名,value是账户外部ID,display是账户名
	ExtendField SCIMExtendField `json:"extendField"` // 用于存放扩展字段的对象	如有扩展字段，则该属性必填
}

// 账户
type SCIMAccount struct {
	UserName     string            `json:"userName"`     // >=4且<18 云IDaaS平台主账户	必填
	PassWord     string            `json:"password"`     // >=6      云IDaaS平台主账户密码	必填
	DisplayName  string            `json:"displayName"`  // >2且<18  用户的显示名称
	Id           string            `json:"id"`           // 用户的唯一id	必填
	Emails       []SCIMEmail       `json:"emails"`       // 邮箱 	和邮箱必有一个
	PhoneNumbers []SCIMPhoneNumber `json:"phoneNumbers"` // 手机号, 只能一个且唯一  和邮箱必有一个
	ExternalId   string            `json:"externalId"`   // 和用户id一样，因为同步过来的，IPG中称为外部id  非必填
	Belongs      []SCIMBelong      `json:"belongs"`      // 所属ou，必须存在  必填
	Locked       bool              `json:"locked"`       // 是否禁用账户，ture禁用账户,false启用账户。禁用账户后将不能登录IDP   非必填，不填默认启用账户
	ExtendField  SCIMExtendField   `json:"extendField"`  // 用于存放扩展字段的对象	如有扩展字段，则该属性必填
}

// 组成员
type SCIMMembers struct {
	Value   string `json:"value"`   // value是账户外部ID
	Display string `json:"display"` // display是账户名
}

// 手机号
type SCIMPhoneNumber struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// 邮箱
type SCIMEmail struct {
	//Primary bool   `json:"primary"`
	Primary string `json:"primary"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

// 所属ou
type SCIMBelong struct {
	BelongOuUuid string `json:"belongOuUuid"`
	OuDirectory  string `json:"ouDirectory"`
	RootNode     bool   `json:"rootNode"`
}

// 管理者
type SCIMManager struct {
	Value       string `json:"value"`       // value代表用户的外部ID,唯一
	DisplayName string `json:"displayName"` // displayName代表用户名
}

// 扩展字段
type SCIMExtendField struct {
	Created     string                 `json:"created"`
	Schemas     []string               `json:"schemas"`
	Description string                 `json:"description"`
	ExpireTime  string                 `json:"expireTime"`
	Attributes  map[string]interface{} `json:"attributes"` // 自定义扩展的字段	如果自定义扩展的字段中的必填选项，则该属性必填
}
