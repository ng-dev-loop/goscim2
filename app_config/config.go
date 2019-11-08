package app_config

import (
	"github.com/astaxie/beego"
)

var (
	KeyUsername = "key_username"
	KeyPassword = "key_password"

	KeyClientId     = "key_client_id"
	KeyClientSecret = "key_client_secret"

	KeyDriverName     = "key_driver_name"
	KeyDataSourceName = "key_data_source_name"
)

func InitConfig() error {

	// BasicAuthenticator验证
	if len(beego.AppConfig.String(KeyUsername)) == 0 {
		_ = beego.AppConfig.Set(KeyUsername, "user")
	}

	if len(beego.AppConfig.String(KeyPassword)) == 0 {
		_ = beego.AppConfig.Set(KeyPassword, "123456")
	}

	if len(beego.AppConfig.String(KeyClientId)) == 0 {
		_ = beego.AppConfig.Set(KeyClientId, "key client id")
	}

	if len(beego.AppConfig.String(KeyClientSecret)) == 0 {
		_ = beego.AppConfig.Set(KeyClientSecret, "key client secret")
	}

	if len(beego.AppConfig.String(KeyDriverName)) == 0 {
		_ = beego.AppConfig.Set(KeyDriverName, "mysql")
	}

	if len(beego.AppConfig.String(KeyDataSourceName)) == 0 {
		_ = beego.AppConfig.Set(KeyDataSourceName, "")
	}

	return nil
}
