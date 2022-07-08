package main

import (
	"github.com/kolesokrut/WheelspinnerBot/internal"
	"github.com/kolesokrut/WheelspinnerBot/internal/config"
	"log"
)

func main() {
	cfg, _ := config.LoadConfig("dev")

	app, err := internal.NewApp(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
