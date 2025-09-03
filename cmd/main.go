package main

import (
	"fmt"
	"log"

	"github.com/jdbdev/liquidity/config"
	"github.com/jdbdev/liquidity/internal/cli"
	"github.com/jdbdev/liquidity/internal/ticker"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Liquidity Pool Manager")

	// Load config settings from .env file
	app := Init()

	//Call ticker service
	tickerService := ticker.NewTickerService(app)

	//Call CLI functions (cmd/cli.go)
	cli.AddPosition()
	cli.ListPositions()
	cli.RemovePosition()
	cli.ShowStatus()
}

func Init() *config.AppConfig {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	app := config.NewAppConfig()
	return app
}
