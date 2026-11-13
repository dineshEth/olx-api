package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dineshEth/olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("usage: migrate <up|down>")
		return
	}

	cfg := config.MostLoad()

	m, err := migrate.New("file://migrations", cfg.Database_URL)
	if err != nil {
		fmt.Println("Error creating migrate instance:", err)
		return
	}

	switch args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal("migrate up:", err)
		}
		log.Println("migrate up successfull")
	case "down":
		if err := m.Steps(-1); err != nil {
			log.Fatal("migrate down:", err)
		}
		log.Println("migrate down successfull")
	default:
		fmt.Println("Invalid command. Use 'up' or 'down'.")
	}
}
