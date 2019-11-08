

适配九州云腾SSO的SCIM服务器

```
go get github.com/astaxie/beego
go get github.com/Go-SQL-Driver/MySQL
go get github.com/satori/go.uuid
go get github.com/go-xorm/xorm
go get github.com/ng-dev-loop/goscim2
```

配置

1. 同步数据库地址
```
key_data_source_name="root:123456@tcp(127.0.0.1:9306)/dbname?charset=utf8"
```

2. 配置九州云腾APP的API Key和API Scret
```
key_client_id = xxxxxxx
key_client_secret = xxxxxxx
```

3. 配置goscim2服务器的验证信息, 该信息用于验证SCIM客户端(basic认证)
```
key_username=root
key_password=123456
```
