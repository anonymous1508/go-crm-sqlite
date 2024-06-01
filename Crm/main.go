package main

import (
	"fmt"
	"log"

	"github.com/anonymous1508/database"
	"github.com/anonymous1508/lead"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	fmt.Println("Connection opened to database")
	if err := database.DBConn.AutoMigrate(&lead.Lead{}); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("failed to start server: ", err)
	}
	// No need to close database connection as gorm v2 handles it.
}
