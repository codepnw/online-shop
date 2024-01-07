package main

import (
	"fmt"

	"github.com/codepnw/online-shop/apps/auth"
	"github.com/codepnw/online-shop/external/database"
	"github.com/codepnw/online-shop/internal/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		return
	}

	if db != nil {
		fmt.Println("database connected..")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	auth.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
