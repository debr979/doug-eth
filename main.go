package main

import (
	"doug/dao"
	"doug/routes"
	"doug/utils/ether"
	lg "doug/utils/logging"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	//load configs
	if err := godotenv.Load("config/config.env"); err != nil {
		lg.Logging.ErrorLogging(err)
	}

	ether.Eth.ClientInit()
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dao.DBSetting.DBSetting(host, dbName, user, password)
}

func main() {

	if err := routes.Routes.Server(); err != nil {
		lg.Logging.ErrorLogging(err)
	}

}
