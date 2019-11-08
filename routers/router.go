// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/auth"
	"github.com/ng-dev/goscim2/app_config"
	"github.com/ng-dev/goscim2/controllers"
	uuid "github.com/satori/go.uuid"
)

func init() {
	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
		beego.NSNamespace("/groups",
			beego.NSInclude(
				&controllers.GroupController{},
			),
		),
		beego.NSNamespace("/organizations",
			beego.NSInclude(
				&controllers.OrganizationController{},
			),
		),
		beego.NSNamespace("/password",
			beego.NSInclude(
				&controllers.PasswordController{},
			),
		),
		beego.NSNamespace("/sso",
			beego.NSInclude(
				&controllers.SSOController{},
			),
		),
	)

	beego.AddNamespace(ns)

	beego.Get("/", func(c *context.Context) {
		c.Redirect(302, "/swagger/")
	})

	// 过滤器
	beego.InsertFilter("/v2/*", beego.BeforeExec, func(context *context.Context) {
		fmt.Printf("recv: %v\n", string(context.Input.RequestBody))
	})

	defaultRealm := uuid.Must(uuid.NewV4()).String()
	beego.InsertFilter("/v2/*", beego.BeforeRouter, auth.NewBasicAuthenticator(

		func(user, pass string) bool {
			username := beego.AppConfig.String(app_config.KeyUsername)
			password := beego.AppConfig.String(app_config.KeyPassword)

			if user == username && pass == password {
				return true
			} else {
				logs.GetBeeLogger().Warn("error auth: %v %v", user, pass)
				return false
			}
		}, defaultRealm))
}
