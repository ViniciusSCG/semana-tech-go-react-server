package main

import (
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    cmd := exec.Command(
        "tern",
        "migrate",
        "--migrations",
        "./internal/store/pgstore/migrations",
        "--config",
        "./internal/store/pgstore/migrations/tern.conf",
    )

    log.Printf("Running command: %s", cmd.String())

    if err := cmd.Run(); err != nil {
        log.Fatalf("Error running command: %v", err)
    }
}