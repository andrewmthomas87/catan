package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	viper.SetConfigName("catan")
	viper.AddConfigPath("/home/config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var (
		zl  *zap.Logger
		err error
	)
	if viper.GetBool("dev") {
		zl, err = zap.NewDevelopment()
	} else {
		zl, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatal(err)
	}
	defer zl.Sync()
	logger := zl.Sugar()

	app := fiber.New()

	addr := viper.GetString("user_server.addr")

	logger.Infow("Starting server", "addr", addr)

	app.Listen(addr)
}
