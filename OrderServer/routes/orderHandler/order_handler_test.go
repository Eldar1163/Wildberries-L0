package orderHandler

import (
	"OrderServer/cache"
	"OrderServer/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"io"
	"log"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

var (
	orderJson, _ = os.ReadFile("../../nats-streaming/resources/correctOrder.json")
)

type MockCache struct {
	order model.Order
}

func (mc *MockCache) ReadOrder(uid string) (model.Order, bool) {
	if uid == "dapo" {
		return mc.order, true
	} else {
		return model.Order{}, false
	}
}

func (mc *MockCache) SaveOrder(order *model.Order) {
}

func (mc *MockCache) InitCache() {
}

func TestGetOrder(t *testing.T) {
	app := fiber.New(fiber.Config{Views: html.New("../../view", ".html")})
	app.Static("static", "../../static")
	app.Post("/order", GetOrder)

	var order model.Order
	err := json.Unmarshal(orderJson, &order)
	if err != nil {
		t.Fatal("Cannot parse JSON")
	}
	cacheInst := MockCache{
		order: order,
	}

	cache.Instance = &cacheInst

	TestTable := []struct {
		Name      string
		OrderUID  string
		isPresent bool
	}{
		{
			Name:      "OrderUID is present",
			OrderUID:  "dapo",
			isPresent: true,
		},
		{
			Name:      "OrderUID is not present",
			OrderUID:  "notpresent",
			isPresent: false,
		},
	}

	for _, subtest := range TestTable {
		t.Run(subtest.Name, func(t *testing.T) {
			form := url.Values{}
			form.Add("uid", subtest.OrderUID)
			req := httptest.NewRequest("POST", "/order", strings.NewReader(form.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.PostForm = form

			resp, err := app.Test(req, -1)

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)

			if subtest.isPresent && !strings.Contains(bodyString, "b563feb7b2b84b6test") ||
				!subtest.isPresent && !strings.Contains(bodyString, "Not found") {
				t.Fatal("Wrong response from server")
			}
		})
	}
}
