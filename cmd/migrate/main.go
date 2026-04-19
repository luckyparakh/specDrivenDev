// cmd/migrate runs all pending database migrations against db.sqlite.
// Usage: go run ./cmd/migrate [up|down]
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	flag.Parse()
	direction := "up"
	if flag.NArg() > 0 {
		direction = flag.Arg(0)
	}

	m, err := migrate.New(
		"file://db/migrations",
		"sqlite://db.sqlite",
	)
	if err != nil {
		log.Fatalf("migrate: init failed: %v", err)
	}
	defer m.Close()

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migrate up: %v", err)
		}
		fmt.Fprintln(os.Stdout, "migrate: up complete")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("migrate down: %v", err)
		}
		fmt.Fprintln(os.Stdout, "migrate: down complete")
	default:
		log.Fatalf("migrate: unknown direction %q (use 'up' or 'down')", direction)
	}
}
