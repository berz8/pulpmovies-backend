package main

import (
	"log"
	"os"

	"github.com/berz8/pulpmovies-backend/database"
	"github.com/berz8/pulpmovies-backend/database/migrations"
	"github.com/berz8/pulpmovies-backend/handlers"
	"github.com/berz8/pulpmovies-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

type (
    GlobalErrorHandlerResp struct {
        Success bool   `json:"success"`
        Message string `json:"message"`
    }
)

func main() {
	// Reading env vars
  isProd := false
  if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
    isProd = true
  } else {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
  }
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	app := fiber.New(fiber.Config{
		Prefork: isProd,
    ErrorHandler: func(c *fiber.Ctx, err error) error {
            return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
                Success: false,
                Message: err.Error(),
            })
    },
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())

	// DB Connection
	database.Connect()
	defer database.DB.Close()

	// Running DB Migrations
  err := migrations.RunMigrations(database.DB)
	if err != nil {
		log.Fatal("error running migrations", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Pulpmovies - Backend API Service")
	})

  routes.AuthRoutes(app)
  routes.UserRoutes(app)

	app.Use(handlers.NotFound)

	log.Fatal(app.Listen(":" + port))
}
