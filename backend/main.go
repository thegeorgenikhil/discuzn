package main

import (
	log "github.com/sirupsen/logrus"

	appLogger "github.com/discuzn/backend/logger"

	"github.com/discuzn/backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config := config.NewAppConfig()
	config.LoadConfig()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	appLogger.InitLogger()

	log.Println("Application started on port: " + config.GetPort())
	err := app.Listen(":" + config.GetPort())

	if err != nil {
		log.Fatal("Error while starting the server: ", err.Error())
	}
}
