package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type dbSetting struct {
	HOST     string `json:"host"`
	DBNAME   string `json:"DBNAME"`
	USER     string `json:"USER"`
	PASSWORD string `json:"PASSWORD"`
}

var DBSetting dbSetting

func (r *dbSetting) DBConnectionSetting() *gorm.DB {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", r.USER, r.PASSWORD, r.HOST, r.DBNAME)
	open, _ := gorm.Open(mysql.Open(connectString), &gorm.Config{})
	return open
}

func (r *dbSetting) DbConn() *gorm.DB {
	return r.DBConnectionSetting()
}

func (r *dbSetting) DBSetting(host, dbName, user, password string) {
	r.HOST = host
	r.DBNAME = dbName
	r.USER = user
	r.PASSWORD = password

	log.Printf("%+v", *r)
}
