package nats_streaming

import (
	"encoding/json"
	"log"

	"github.com/noglezneva/wbl0/internal/model"
	"github.com/noglezneva/wbl0/internal/service/db"

	"github.com/nats-io/stan.go"
)

const subj = "order"
const natsURL = "http://nats:4222"

// Функция обработчика подтверждений
func ackHandler(ackedID string, err error) {
	if err != nil {
		log.Printf("ERROR publishing msg id %s: %v\n", ackedID, err.Error())
	}
}

// Функция обработчика полученных сообщений
func (sc Connect) ackSubscribe() func(m *stan.Msg) {
	return func(m *stan.Msg) {
		log.Printf("received a message! %s\n", m.Data)
		order := new(model.Order)
		err := json.Unmarshal(m.Data, order)
		if err != nil {
			log.Println("unmarshal error")
			log.Println(err)
			return
		}
		err = sc.dbService.Create(order)
		if err != nil {
			log.Println("db error")
			return
		}
	}
}

type Connect struct {
	sc        stan.Conn
	dbService *db.Service
}

// Функция создания нового соединения
func NewConn(dbService *db.Service) *Connect {
	sc, err := stan.Connect("test-cluster", "test-cluster", stan.NatsURL(natsURL))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return &Connect{sc: sc, dbService: dbService}
}

// Метод публикации заказа
func (sc Connect) PublishOrder(order *model.Order) {
	// Преобразуем заказ в JSON
	mes, err := json.Marshal(*order)
	if err != nil {
		log.Println(err)
		return
	}
	id, err := sc.sc.PublishAsync(subj, mes, ackHandler) // returns immediately
	if err != nil {
		log.Printf("error publishing msg %s: %v\n", id, err.Error())
	}
}

// Метод подписки на заказы
func (sc Connect) SubscribeOrder() stan.Subscription {
	var err error
	sub, err := sc.sc.Subscribe(
		subj,
		sc.ackSubscribe(),
		stan.DurableName(subj))
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("subscribed to subject\n")
	return sub
}
