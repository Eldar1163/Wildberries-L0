package model

type Item struct {
	Id          int `gorm:"primaryKey"`
	OrderId     string
	ChrtId      int     `json:"chrt_id" validate:"min=0"`
	TrackNumber string  `json:"track_number" validate:"required"`
	Price       float64 `json:"price" validate:"min=0"`
	Rid         string  `json:"rid" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Sale        int     `json:"sale" validate:"min=0"`
	Size        string  `json:"size" validate:"required"`
	TotalPrice  float64 `json:"total_price" validate:"min=0"`
	NmId        int     `json:"nm_id" validate:"min=0"`
	Brand       string  `json:"brand" validate:"required"`
	Status      int     `json:"status" validate:"min=0"`
}
