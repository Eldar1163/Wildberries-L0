package nats_streaming

import (
	"OrderServer/model"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	"os"
	"testing"
)

var (
	correctOrder, _   = os.ReadFile("./resources/correctOrder.json")
	corruptedOrder, _ = os.ReadFile("./resources/corruptedOrder.json")
	notValidOrder, _  = os.ReadFile("./resources/notValidOrder.json")
)

type MockOrderCreator struct {
	calls int
}

func (moc *MockOrderCreator) CreateOrder(order *model.Order) {
	moc.calls++
}

func TestNatsStreamingSubscribe(t *testing.T) {
	NatsTest := []struct {
		name                     string
		msg                      stan.Msg
		orderBytes               []byte
		mom                      MockOrderCreator
		expectedCreateOrderCalls int
	}{
		{
			name: "SuccessCreateOrder",
			msg: stan.Msg{
				MsgProto: pb.MsgProto{
					Data: correctOrder,
				},
				Sub: nil,
			},
			mom:                      MockOrderCreator{},
			orderBytes:               correctOrder,
			expectedCreateOrderCalls: 1,
		},

		{
			name: "FailedCorruptedOrder",
			msg: stan.Msg{
				MsgProto: pb.MsgProto{
					Data: corruptedOrder,
				},
				Sub: nil,
			},
			mom:                      MockOrderCreator{},
			orderBytes:               corruptedOrder,
			expectedCreateOrderCalls: 0,
		},

		{
			name: "FailedNotValidOrder",
			msg: stan.Msg{
				MsgProto: pb.MsgProto{
					Data: notValidOrder,
				},
				Sub: nil,
			},
			mom:                      MockOrderCreator{},
			orderBytes:               notValidOrder,
			expectedCreateOrderCalls: 0,
		},
	}

	for _, subtest := range NatsTest {
		t.Run(subtest.name, func(t *testing.T) {
			order := model.Order{}
			_ = json.Unmarshal(subtest.orderBytes, &order)
			orderCreator = &subtest.mom
			handleOrder(&subtest.msg)
			if subtest.mom.calls != subtest.expectedCreateOrderCalls {
				t.Error("Correct Order data was not stored")
			}
		})
	}
}
