package main

import (
	"SubscriptionService/internal/config"
	"SubscriptionService/internal/database"
	"SubscriptionService/internal/migration"
	"SubscriptionService/internal/subscription/handlers"
	"SubscriptionService/internal/subscription/repository"
	"SubscriptionService/internal/subscription/service"
	"SubscriptionService/pkg"
	"log/slog"
)

func main() {
	logger := pkg.SetupLogger()

	cfg := config.Load()

	db, err := database.InitDB(cfg)
	if err != nil {
		slog.Error("Error on openning DB - ", "Err: ", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			slog.Error("Error while closing DB - ", "Err: ", err)
		}
	}()

	migration.RunMigrations(db)

	rep := repository.InitRepository(db.DB)
	service := service.InitService(rep, logger)
	handler := handlers.InitHandler(service, logger)

	router := handlers.NewRouter(handler)

	router.Run(":8080")
}