package models

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
)

var (
	DBEngine *xorm.Engine
)

/**
driverName := "mysql"
dataSourceName := "root:123456@tcp(localhost:23306)/test3?charset=utf8"
showSQL := true

url: [username]:[password]@tcp([ip]:[port])/[database]?charset=utf8
*/
func InitDataBase(driverName string, dataSourceName string, showSQL bool) error {
	var err error

	DBEngine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		logs.GetBeeLogger().Error("new DBEngine fail %v", err)
		return err
	} else {
		logs.GetBeeLogger().Info("database type: %v", DBEngine.Dialect().DBType())
	}

	DBEngine.ShowSQL(showSQL)
	DBEngine.SetMaxIdleConns(5)
	DBEngine.SetMaxOpenConns(5)
	//DBEngine.StoreEngine("MyISAM") //InnoDB

	/*s := DBEngine.NewSession()
	result, err := s.Query("select name from userinfo;")
	if err != nil {
		logs.GetBeeLogger().Error("%v", err)
	} else {
		for _, obj := range result {
			for key, val := range obj {
				fmt.Printf("key: %v, val: %v\n", key, string(val))
			}
		}
	}*/

	return nil
}
