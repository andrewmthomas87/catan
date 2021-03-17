package main

import (
	"context"
	"fmt"
	"log"

	"github.com/andrewmthomas87/catan/ent"
	_ "github.com/lib/pq"
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

	logger.Info("Opening client connection...")

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", viper.GetString("postgres.host"), viper.GetString("postgres.port"), viper.GetString("postgres.user"), viper.GetString("postgres.dbname"), viper.GetString("postgres.password")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	logger.Info("Client connection opened")
	logger.Info("Running auto-migration...")

	if err := client.Schema.Create(context.Background()); err != nil {
		logger.Fatalf("failed creating schema resources: %v", err)
	}

	logger.Info("Auto-migration completed")
}
