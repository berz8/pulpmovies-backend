package main

import (
	"log"
	"os"

	"github.com/berz8/pulpmovies-backend/database"
	"github.com/berz8/pulpmovies-backend/database/migrations"
	"github.com/berz8/pulpmovies-backend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Reading env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	isProd := false
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		isProd = true
	}

	app := fiber.New(fiber.Config{
		Prefork: isProd,
	})

  app.Use(logger.New())
  app.Use(recover.New())
  app.Use(compress.New())


  // DB Connection
  database.Connect()
  defer database.DB.Close()

  // Running DB Migrations
  err = migrations.RunMigrations(database.DB)
  if err != nil {
    log.Fatal("error running migrations", err)
  }
  

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Pulpmovies - Backend API Service")
	})

  app.Use(handlers.NotFound)

	log.Fatal(app.Listen(":" + port))
}
