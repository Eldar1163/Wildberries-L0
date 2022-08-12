package model

type Payment struct {
	OrderId      string  `gorm:"primaryKey"`
	Transaction  string  `json:"transaction" validate:"required"`
	RequestId    *string `json:"request_id"`
	Currency     string  `json:"currency" validate:"required"`
	Provider     string  `json:"provider" validate:"required"`
	Amount       float64 `json:"amount" validate:"required"`
	PaymentDt    int     `json:"payment_dt" validate:"required"`
	Bank         string  `json:"bank" validate:"required"`
	DeliveryCost float64 `json:"delivery_cost" validate:"required"`
	GoodsTotal   int     `json:"goods_total" validate:"min=0"`
	CustomFee    float64 `json:"custom_fee" validate:"min=0"`
}
