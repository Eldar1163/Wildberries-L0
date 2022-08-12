package router

import (
	"OrderServer/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("", logger.New())
	routes.SetupOrderRoutes(api)
}
