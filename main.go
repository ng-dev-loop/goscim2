package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ng-dev/goscim2/app_config"
	"github.com/ng-dev/goscim2/models"
	_ "github.com/ng-dev/goscim2/routers"
)

func main() {
	err := app_config.InitConfig()
	if err != nil {
		logs.GetBeeLogger().Error("init config fail %v", err)
		return
	}

	_ = models.InitDataBase(
		beego.AppConfig.String(app_config.KeyDriverName),
		beego.AppConfig.String(app_config.KeyDataSourceName),
		false)

	//if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	beego.BConfig.CopyRequestBody = true

	beego.Run()
}
