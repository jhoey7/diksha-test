package main

import (
	"diskha-test/models"
	_ "diskha-test/router"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os/user"
	"path/filepath"
)

const (
	keyDbUser     = "DBUsername"
	keyDbPassword = "DBPassword"
	keyDbHost     = "DBHost"
	keyDbPort     = "DBPort"
	keyDbName     = "DBName"
)

func init() {
	cfglog := `{"filename":"logs/app.log","level":7,"daily":true,"maxdays":365,"rotate":true}`
	beego.SetLogger("file", cfglog)

	dbUser, dbPassword, dbHost, dbPort, dbName := getDatabaseConfig()
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		panic(err.Error())
	}

	orm.Debug = beego.AppConfig.DefaultBool("orm.query.debug", true)
	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Warehouse))
	orm.RegisterModel(new(models.Products))
	orm.RegisterModel(new(models.ProductDetails))
	orm.RegisterModel(new(models.Orders))
	orm.RegisterModel(new(models.OrdersDetail))
}

func getDatabaseConfig() (string, string, string, string, string) {
	dbUser := beego.AppConfig.String(keyDbUser)
	dbPassword := beego.AppConfig.String(keyDbPassword)
	dbHost := beego.AppConfig.String(keyDbHost)
	dbPort := beego.AppConfig.String(keyDbPort)
	dbName := beego.AppConfig.String(keyDbName)
	usr, err := user.Current()
	if err == nil {
		configPath := filepath.Join(usr.HomeDir, "conf", "dealls-prototype.conf")
		conf, err := config.NewConfig("ini", configPath)
		if err == nil {
			log.Println("External configuration found in ", configPath)
			dbUser = conf.DefaultString(keyDbUser, beego.AppConfig.String(keyDbUser))
			dbPassword = conf.DefaultString(keyDbPassword, beego.AppConfig.String(keyDbPassword))
			dbHost = conf.DefaultString(keyDbHost, beego.AppConfig.String(keyDbHost))
			dbPort = conf.DefaultString(keyDbPort, beego.AppConfig.String(keyDbPort))
			dbName = conf.DefaultString(keyDbName, beego.AppConfig.String(keyDbName))
		}
	}
	return dbUser, dbPassword, dbHost, dbPort, dbName
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
