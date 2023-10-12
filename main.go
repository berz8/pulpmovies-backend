package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Pulpmovies - Backend API Service")
	})

	app.Listen(":" + port)

}
