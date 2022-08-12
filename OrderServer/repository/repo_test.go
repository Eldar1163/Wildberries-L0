package repository

import (
	"OrderServer/cache"
	"OrderServer/model"
	"OrderServer/repository/db"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"regexp"
	"strconv"
	"testing"
)

var (
	orderJSON, _ = os.ReadFile("../nats-streaming/resources/correctOrder.json")
)

type MockCache struct {
	saveOrderCalls int
}

func (mc *MockCache) ReadOrder(uid string) (model.Order, bool) {
	return model.Order{}, false
}

func (mc *MockCache) SaveOrder(order *model.Order) {
	mc.saveOrderCalls++
}

func (mc *MockCache) InitCache() {
}

func TestCreateOrder(t *testing.T) {
	TestTable := []struct {
		Name           string
		OrderIdCnt     int
		TrackNumberCnt int
		OrderCreate    bool
		MockCache      MockCache
	}{
		{
			Name:           "Success create",
			OrderIdCnt:     0,
			TrackNumberCnt: 0,
			OrderCreate:    true,
			MockCache:      MockCache{},
		},
		{
			Name:           "Order UID is present",
			OrderIdCnt:     1,
			TrackNumberCnt: 0,
			OrderCreate:    false,
			MockCache:      MockCache{},
		},
		{
			Name:           "Track number is present",
			OrderIdCnt:     0,
			TrackNumberCnt: 1,
			OrderCreate:    false,
			MockCache:      MockCache{},
		},
		{
			Name:           "Order UID and track number is present",
			OrderIdCnt:     1,
			TrackNumberCnt: 1,
			OrderCreate:    false,
			MockCache:      MockCache{},
		},
	}

	for _, subtest := range TestTable {
		t.Run(subtest.Name, func(t *testing.T) {
			var order = model.Order{}
			err := json.Unmarshal(orderJSON, &order)
			if err != nil {
				t.Fatal("Cannot parse correct json")
			}

			datab, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer datab.Close()

			mock.MatchExpectationsInOrder(false)
			mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM \"order\" WHERE order_uid = $1")).
				WithArgs(order.OrderUid).
				WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(strconv.Itoa(subtest.OrderIdCnt)))
			if subtest.OrderIdCnt == 0 {
				mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM \"order\" WHERE track_number = $1")).
					WithArgs(order.TrackNumber).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(strconv.Itoa(subtest.TrackNumberCnt)))
			}

			dialector := postgres.New(postgres.Config{
				DSN:                  "sqlmock_db_0",
				DriverName:           "postgres",
				Conn:                 datab,
				PreferSimpleProtocol: true,
			})
			db.DB, err = gorm.Open(dialector, &gorm.Config{})

			cache.Instance = &subtest.MockCache
			(&OrderCreatorImpl{}).CreateOrder(&order)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

			if subtest.OrderCreate {
				if subtest.MockCache.saveOrderCalls == 0 {
					t.Error("Order was not stored in cache")
				}
			} else {
				if subtest.MockCache.saveOrderCalls == 1 {
					t.Error("Order was stored in cache")
				}
			}
		})
	}
}
