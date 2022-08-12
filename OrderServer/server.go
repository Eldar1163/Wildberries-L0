package main

import (
	"OrderServer/cache"
	"OrderServer/logger"
	nats_streaming "OrderServer/nats-streaming"
	"OrderServer/repository/db"
	"OrderServer/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	db.DatabasebOpen()
	cache.Instance.InitCache()
	app := fiber.New(fiber.Config{Views: html.New("./view", ".html")})
	app.Static("static", "./static")
	router.SetupRoutes(app)
	nats_streaming.NatsStreamingSetup()
	err := app.Listen(":8080")
	if err != nil {
		logger.ErrorLogger.Println("Cannot listen http requests")
	}
}
