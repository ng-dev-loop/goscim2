package controllers

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/ng-dev/goscim2/app_config"
	"io/ioutil"
	"net/http"
)

type JsonHttpResult struct {
	ErrorNumber int      `json:"errorNumber"`
	Errors      []string `json:"errors"`
}

type SSOController struct {
	beego.Controller
}

type TicketInfo struct {
	UserName string `json:"username"`
	Error    int    `json:"error"`
	ErrorMsg string `json:"errorMsg"`
}

// @Title Get
// @Description find object by groupId
// @Param	objectId		path 	string	true		"the groupId you want to get"
// @Success 200 {object} models.SCIMGroupModel
// @Failure 403 :groupId is empty
// @router / [get]
func (o *SSOController) Get() {
	ticket := o.GetString("code")
	logs.GetBeeLogger().Info("ticket: %v", ticket)

	// 2.访问IPG平台去验证此次的ticket是否合法
	// {IPGServer} 为你的IPG服务器地址，例 https://jzyt.idsmanager.com
	// {applicationId} 为你创建的CAS应用的id
	// String url = "http://{IPGServer}/public/api/application/cas/callback_{applicationId}?code=" + ticket;
	IPGServer := "hytera.idsmanager.com"
	applicationId := "hyteracas"

	url := fmt.Sprintf("http://%v/public/api/application/cas/callback_%v?code=%v", IPGServer, applicationId, ticket)
	response, err := http.Get(url)
	if err != nil {
		logs.GetBeeLogger().Error("get ticket info fail %v", err)
		o.Ctx.WriteString(fmt.Sprintf("get ticket info fail %v", err))
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		o.Ctx.WriteString(fmt.Sprintf("read ticket info fail %v", err))
		logs.GetBeeLogger().Error("read ticket info fail %v", err)
	} else {
		ticketInfo := string(data)
		logs.GetBeeLogger().Info(ticketInfo)
		if response.StatusCode == http.StatusOK {

		}

		var objTicket TicketInfo
		err = json.Unmarshal(data, &objTicket)
		if err != nil {
			o.Ctx.WriteString(fmt.Sprintf("ticket info unmarshal fail %v", err))
			logs.GetBeeLogger().Error("ticket info unmarshal fail %v", err)
			return
		}

		if objTicket.Error == 0 {
			o.Redirect("https://www.baidu.com", http.StatusFound)
		} else {
			o.Ctx.WriteString(fmt.Sprintf("hello error ticket: %v", ticketInfo))
		}
	}
}

// @router /login [get]post login fail
func (o *SSOController) Login() {
	type CASParam struct {
		ClientId     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
		Username     string `json:"username"`
		Password     string `json:"password"`
	}

	param := CASParam{
		ClientId:     beego.AppConfig.String(app_config.KeyClientId),
		ClientSecret: beego.AppConfig.String(app_config.KeyClientSecret),
		Username:     o.GetString("username", ""),
		Password:     o.GetString("password", ""),
	}
	data, _ := json.Marshal(&param)

	// 忽略证书检查
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{Transport: tr}

	// 发送请求
	url := "https://hytera.idsmanager.com/public/enduser/login"
	res, err := client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		o.Ctx.Output.SetStatus(http.StatusInternalServerError)
		logs.GetBeeLogger().Error("post login fail %v", err)
		fmt.Println(err)
		return
	}

	// 读取回复
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		o.Ctx.Output.SetStatus(http.StatusInternalServerError)
		logs.GetBeeLogger().Error("read response fail %v", err)
		return
	}
	_ = res.Body.Close()

	// 回复
	o.Ctx.Output.SetStatus(res.StatusCode)
	for key, value := range res.Header {
		for _, v := range value {
			o.Ctx.Output.Header(key, v)
		}
	}
	_ = o.Ctx.Output.Body(data)
}
