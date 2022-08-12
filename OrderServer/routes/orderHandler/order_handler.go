package orderHandler

import (
	"OrderServer/cache"
	"OrderServer/dto"
	"OrderServer/dto/mappers"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func MainPage(c *fiber.Ctx) error {
	return c.Render("main", fiber.Map{})
}

func GetOrder(c *fiber.Ctx) error {
	uid := c.FormValue("uid")
	if order, found := cache.Instance.ReadOrder(uid); found {
		orderDTO := dto.Order{}
		mappers.MapOrderToOrderDto(&order, &orderDTO)
		orderPrettyText, _ := json.MarshalIndent(orderDTO, "", "    ")
		return c.Render("main", fiber.Map{"order": string(orderPrettyText)})
	} else {
		return c.Render("notfound", fiber.Map{})
	}
}
