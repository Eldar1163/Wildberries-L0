package nats_streaming

import (
	"OrderServer/logger"
	"OrderServer/model"
	"OrderServer/repository"
	"encoding/json"
	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "order-subscriber"
	channel   = "order-notification"
)

var (
	SC           stan.Conn
	orderCreator repository.OrderCreator = &repository.OrderCreatorImpl{}
)

func handleOrder(orderMSG *stan.Msg) {
	order := model.Order{}

	err := json.Unmarshal(orderMSG.Data, &order)
	if err != nil {
		logger.WarningLogger.Println("Cannot unmarshal data from nats-streaming-server")
		return
	}
	if err = order.Validate(); err != nil {
		logger.WarningLogger.Println("Cannot validate incoming json\n" + err.Error())
		return
	}

	orderCreator.CreateOrder(&order)
}

func natsStreamingSubscribe() {
	_, err := SC.Subscribe(channel, handleOrder, stan.StartWithLastReceived())
	if err != nil {
		logger.ErrorLogger.Panic("Cannot subscribe to nats-streaming-server chanel")
	}
}

func NatsStreamingSetup() {
	var err error
	SC, err = stan.Connect(clusterID, clientID)
	if err != nil {
		logger.ErrorLogger.Panic("Cannot connect to nats-streaming-server")
	}
	natsStreamingSubscribe()
}
