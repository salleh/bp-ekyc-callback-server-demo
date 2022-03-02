package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salleh/bp-ekyc-callback-server-demo/config"
	"github.com/salleh/bp-ekyc-callback-server-demo/router"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	if config.Debug() {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "BP eKYC Callback Demo Server v0.0.1",
	})

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
