package main

import (
	"doug/routes"
	"doug/utils/ether"
	lg "doug/utils/logging"
	"github.com/joho/godotenv"
)

func init() {
	//load configs
	if err := godotenv.Load("config/config.env"); err != nil {
		lg.Logging.ErrorLogging(err)
	}

	ether.Eth.ClientInit()
}

func main() {
	if err := routes.Routes.Server(); err != nil {
		lg.Logging.ErrorLogging(err)
	}
}
