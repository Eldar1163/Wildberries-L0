package model

import (
	"errors"
	"time"
)

type Order struct {
	OrderUid          string    `gorm:"primaryKey" json:"order_uid" validate:"required"`
	Delivery          Delivery  `gorm:"foreignKey:OrderId" json:"delivery" validate:"required"`
	Payment           Payment   `gorm:"foreignKey:OrderId" json:"payment" validate:"required"`
	Items             []Item    `gorm:"foreignKey:OrderId" json:"items" validate:"required"`
	TrackNumber       string    `json:"track_number" validate:"required"`
	Entry             string    `json:"entry" validate:"required"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature *string   `json:"internal_signature"`
	CustomerId        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	SmId              int       `json:"sm_id" validate:"min=0"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}

func (order *Order) Validate() error {
	var errStr string
	if order.InternalSignature == nil {
		errStr += "Field 'InternalSignature' is null\n"
	}
	if order.Payment.RequestId == nil {
		errStr += "Field 'RequestId' is null"
	}
	if errStr == "" {
		return Validator.Struct(order)
	} else {
		return errors.New(errStr)
	}
}
