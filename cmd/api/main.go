package main

import (
	"fmt"

	"github.com/codepnw/online-shop/external/database"
	"github.com/codepnw/online-shop/internal/config"
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
}
