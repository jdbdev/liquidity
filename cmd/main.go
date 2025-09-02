package main

import (
	"fmt"

	"github.com/jdbdev/liquidity/internal/cli"
)

func main() {
	fmt.Println("Liquidity Pool Manager")

	//Call ticker service
	NewTickerService()

	//Call CLI functions (cmd/cli.go)
	cli.AddPosition()
	cli.ListPositions()
	cli.RemovePosition()
	cli.ShowStatus()
}
