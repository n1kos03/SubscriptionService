package main

import (
	"SubscriptionService/internal/config"
	"SubscriptionService/internal/database"
	"SubscriptionService/internal/migration"
	"log/slog"
)

func main() {
	cfg := config.Load()

	db, err := database.InitDB(cfg)
	if err != nil {
		slog.Error("Error on openning DB - ", "err: ", err)
		return
	}

	migration.RunMigrations(db)
}