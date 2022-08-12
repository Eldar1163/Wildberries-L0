package routes

import (
	"OrderServer/routes/orderHandler"
	"github.com/gofiber/fiber/v2"
)

func SetupOrderRoutes(router fiber.Router) {
	order := router.Group("/order")
	order.Get("/", orderHandler.MainPage)
	order.Post("/", orderHandler.GetOrder)
}
