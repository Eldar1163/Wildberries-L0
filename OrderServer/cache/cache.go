package cache

import (
	"OrderServer/model"
	"OrderServer/repository/db"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm/clause"
	"time"
)

type GoCacheInterface interface {
	SaveOrder(value *model.Order)
	ReadOrder(uid string) (model.Order, bool)
	InitCache()
}

type GoCache struct {
	instance *cache.Cache
}

var (
	Instance GoCacheInterface = &GoCache{}
)

func (c *GoCache) InitCache() {
	c.instance = cache.New(5*time.Minute, 10*time.Minute)
	c.restoreCache()
}

func (c *GoCache) SaveOrder(value *model.Order) {
	c.instance.Set(value.OrderUid, *value, cache.NoExpiration)
}

func (c *GoCache) ReadOrder(uid string) (model.Order, bool) {
	order, found := c.instance.Get(uid)
	if found {
		return order.(model.Order), found
	} else {
		return model.Order{}, found
	}
}

func (c *GoCache) restoreCache() {
	database := db.DB
	var orders []model.Order

	database.
		Preload("Payment").
		Preload("Delivery").
		Preload("Items").
		Preload(clause.Associations).
		Find(&orders)
	for _, order := range orders {
		c.SaveOrder(&order)
	}
}
