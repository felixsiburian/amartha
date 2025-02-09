package main

import (
	"amartha/config"
	"amartha/internal/domain"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	app := config.Config{}

	app.CatchError(app.InitEnv())

	dbConfig := app.GetDBConfig()

	db := config.ConnectionDB(dbConfig)

	AutoMigrate(db)
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&domain.Loan{},
		&domain.Payment{},
	).Error; err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	log.Println("Database migration completed successfully")
}
