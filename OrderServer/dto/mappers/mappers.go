package mappers

import (
	"OrderServer/dto"
	"OrderServer/logger"
	"OrderServer/model"
	"github.com/devfeel/mapper"
)

func MapOrderToOrderDto(order *model.Order, orderDTO *dto.Order) {
	deliveryDto := dto.Delivery{}
	MapDeliveryToDeliveryDto(&order.Delivery, &deliveryDto)

	paymentDto := dto.Payment{}
	MapPaymentToPaymentDto(&order.Payment, &paymentDto)

	var itemsDto []dto.Item
	for ind, item := range order.Items {
		itemsDto = append(itemsDto, dto.Item{})
		MapItemToItemDto(&item, &itemsDto[ind])
	}

	orderDTO.OrderUid = order.OrderUid
	orderDTO.TrackNumber = order.TrackNumber
	orderDTO.Entry = order.Entry
	orderDTO.Delivery = deliveryDto
	orderDTO.Payment = paymentDto
	orderDTO.Items = itemsDto
	orderDTO.Locale = order.Locale
	orderDTO.InternalSignature = order.InternalSignature
	orderDTO.CustomerId = order.CustomerId
	orderDTO.DeliveryService = order.DeliveryService
	orderDTO.Shardkey = order.Shardkey
	orderDTO.SmId = order.SmId
	orderDTO.DateCreated = order.DateCreated
	orderDTO.OofShard = order.OofShard
}

func MapItemToItemDto(item *model.Item, itemDto *dto.Item) {
	err := mapper.Mapper(item, itemDto)
	if err != nil {
		logger.ErrorLogger.Panic("Cannot map Item to Item DTO")
	}
}

func MapDeliveryToDeliveryDto(delivery *model.Delivery, deliveryDto *dto.Delivery) {
	err := mapper.Mapper(delivery, deliveryDto)
	if err != nil {
		logger.ErrorLogger.Panic("Cannot map Delivery to Delivery DTO")
	}
}

func MapPaymentToPaymentDto(payment *model.Payment, paymentDto *dto.Payment) {
	err := mapper.Mapper(payment, paymentDto)
	if err != nil {
		logger.ErrorLogger.Panic("Cannot map Payment to Payment DTO")
	}
}
