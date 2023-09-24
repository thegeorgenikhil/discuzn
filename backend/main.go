package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/discuzn/backend/db"
	appLogger "github.com/discuzn/backend/logger"

	"github.com/discuzn/backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config := config.NewAppConfig()
	config.LoadConfig()

	appLogger.InitLogger()

	fmt.Println(config.GetDbURI())

	conn,err := db.ConnectDB(config.GetDbURI())

	if err != nil {
		log.Println("Error while connecting to DB", err)
	}

	db.DB = conn

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("ok")
	})


	log.Println("Application started on port: " + config.GetPort())
	err = app.Listen(":" + config.GetPort())

	if err != nil {
		log.Fatal("Error while starting the server: ", err.Error())
	}
}
