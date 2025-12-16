package migration

import (
	"SubscriptionService/internal/database"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed sql/*.sql
var embedMigrations embed.FS

func RunMigrations(db *database.Database) {
	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Println("Error in taking sql.DB connection: ", err)
		return
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Println("Error on setup dialect for migrations: ", err)
		return
	}

	if err := goose.Up(sqlDB, "sql"); err != nil {
		log.Println("Error while applying migrations: ", err)
		return
	}
}
