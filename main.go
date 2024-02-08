package main

import (
	"simplink/config"
	"simplink/helpers/env"
	"simplink/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	var mysqlConfig = new(config.Mysql)
	if err := env.Load(nil, mysqlConfig); err != nil {
		panic(err)
	}

	mysqlConnection, err := database.MysqlInit(*mysqlConfig)
	if err != nil {
		panic(err)
	}

	if err := database.MysqlMigrate(mysqlConnection); err != nil {
		panic(err)
	}

	app := echo.New()

	app.Start(":8000")
}
