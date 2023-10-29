package nats_streaming

import (
	"time"

	"github.com/noglezneva/wbl0/internal/model"
	"github.com/noglezneva/wbl0/internal/service/db"
)

func SubscribeAndPublishTest(d *db.Service) {
	stan := NewConn(d)

	// Подписываемся канал заказов
	stan.SubscribeOrder()

	// Создаем два элемента модели Item
	items := make([]model.Item, 2)

	items[0] = model.Item{
		ChrtId:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmId:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}

	items[1] = model.Item{
		ChrtId:      7934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       250,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Face cream",
		Sale:        15,
		Size:        "0",
		TotalPrice:  210,
		NmId:        6565612,
		Brand:       "Clinique",
		Status:      200,
	}
	// Создаем объект модели Delivery
	delivery := model.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}
	// Создаем объект модели Payment
	payment := model.Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestId:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}
	// Создаем объект модели Order
	order := &model.Order{
		OrderUID:          "b563feb7b2b84b6test",
		TrackNumber:       "WBILMTESTTRACK",
		Entry:             "WBIL",
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       time.Time{},
		OofShard:          "1",
	}
	// Отправляем два сообщения о заказе
	stan.PublishOrder(order)
	stan.PublishOrder(order)
}
