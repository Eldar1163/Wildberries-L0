package mappers

import (
	"OrderServer/dto"
	"OrderServer/model"
	"reflect"
	"testing"
	"time"
)

var (
	reqId  = "testID"
	intSig = "intSig"
	order  = model.Order{
		OrderUid: "test",
		Delivery: model.Delivery{
			OrderId: "test",
			Name:    "Kek",
			Phone:   "+79184201896",
			Zip:     "350000",
			City:    "Moscow",
			Address: "Ul. Pushkina, d. 20",
			Region:  "Moscow oblast",
			Email:   "test@mail.ru",
		},
		Payment: model.Payment{
			OrderId:      "test",
			Transaction:  "testtesttest",
			RequestId:    &reqId,
			Currency:     "USD",
			Provider:     "TEST",
			Amount:       350.12,
			PaymentDt:    1637283423,
			Bank:         "Alfa",
			DeliveryCost: 500.12,
			GoodsTotal:   5,
			CustomFee:    12.12,
		},
		Items: []model.Item{
			{
				Id:          1,
				OrderId:     "Test",
				ChrtId:      12,
				TrackNumber: "TESTTRACK",
				Price:       510.50,
				Rid:         "testID",
				Name:        "testname",
				Sale:        54,
				Size:        "XXL",
				TotalPrice:  609.50,
				NmId:        45,
				Brand:       "Apple",
				Status:      200,
			},
		},
		TrackNumber:       "WILDBERRIESTRACK",
		Entry:             "TEST",
		Locale:            "RU",
		InternalSignature: &intSig,
		CustomerId:        "testID",
		DeliveryService:   "WILD",
		Shardkey:          "test",
		SmId:              123,
		DateCreated:       time.Now(),
		OofShard:          "testOOF",
	}

	expectedOrderDTO = dto.Order{
		OrderUid: "test",
		Delivery: dto.Delivery{
			Name:    "Kek",
			Phone:   "+79184201896",
			Zip:     "350000",
			City:    "Moscow",
			Address: "Ul. Pushkina, d. 20",
			Region:  "Moscow oblast",
			Email:   "test@mail.ru",
		},
		Payment: dto.Payment{
			Transaction:  "testtesttest",
			RequestId:    &reqId,
			Currency:     "USD",
			Provider:     "TEST",
			Amount:       350.12,
			PaymentDt:    1637283423,
			Bank:         "Alfa",
			DeliveryCost: 500.12,
			GoodsTotal:   5,
			CustomFee:    12.12,
		},
		Items: []dto.Item{
			{
				ChrtId:      12,
				TrackNumber: "TESTTRACK",
				Price:       510.50,
				Rid:         "testID",
				Name:        "testname",
				Sale:        54,
				Size:        "XXL",
				TotalPrice:  609.50,
				NmId:        45,
				Brand:       "Apple",
				Status:      200,
			},
		},
		TrackNumber:       "WILDBERRIESTRACK",
		Entry:             "TEST",
		Locale:            "RU",
		InternalSignature: &intSig,
		CustomerId:        "testID",
		DeliveryService:   "WILD",
		Shardkey:          "test",
		SmId:              123,
		DateCreated:       time.Now(),
		OofShard:          "testOOF",
	}
)

func TestCorrectMapOrderToOrderDto(t *testing.T) {
	gotOrderDTO := dto.Order{}
	MapOrderToOrderDto(&order, &gotOrderDTO)

	if !reflect.DeepEqual(gotOrderDTO, expectedOrderDTO) {
		t.Error("Expected and got orderDTO are not equal")
	}
}
