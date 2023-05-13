package main

import (
	"go-clean-aws/app/config"
	"go-clean-aws/app/database"
	"go-clean-aws/app/router"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDB(*cfg)
	database.Migrate(db)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(db, e)

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
