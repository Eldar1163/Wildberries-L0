package repository

import (
	"OrderServer/cache"
	"OrderServer/model"
	"OrderServer/repository/db"
)

type OrderCreator interface {
	CreateOrder(order *model.Order)
}

type OrderCreatorImpl struct {
}

func (d *OrderCreatorImpl) CreateOrder(order *model.Order) {
	var cnt1 int64
	var cnt2 int64

	db.DB.Table("\"order\"").Where("order_uid = ?", order.OrderUid).Count(&cnt1)
	if cnt1 != 0 {
		return
	}
	db.DB.Table("\"order\"").Where("track_number = ?", order.TrackNumber).Count(&cnt2)
	if cnt2 != 0 {
		return
	}

	db.DB.Create(order)
	cache.Instance.SaveOrder(order)
}
