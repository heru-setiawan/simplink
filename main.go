package main

import (
	"simplink/config"
	"simplink/features/links/handler"
	"simplink/features/links/repository"
	"simplink/features/links/service"
	"simplink/helpers/env"
	"simplink/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	if err := database.MysqlMigrate(mysqlConnection, &repository.Link{}); err != nil {
		panic(err)
	}

	linkRepository := repository.NewLinkRepository(mysqlConnection)
	linkService := service.NewLinkService(linkRepository)
	linkHandler := handler.NewLinkHandler(linkService)

	app := echo.New()
	app.Use(middleware.Logger())

	app.GET("/:shorten-link", linkHandler.GetByShort)
	app.POST("/links", linkHandler.Create)

	app.Start(":8000")
}
