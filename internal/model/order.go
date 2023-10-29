package model

import "time"

// Order представляет информацию о заказе.
type Order struct {
	ID                int64     // Уникальный идентификатор заказа
	OrderUID          string    `json:"order_uid"`          // Уникальный идентификатор заказа
	TrackNumber       string    `json:"track_number"`       // Номер отслеживания заказа
	Entry             string    `json:"entry"`              // Входные данные заказа
	Delivery          Delivery  `json:"delivery"`           // Информация о доставке
	Payment           Payment   `json:"payment"`            // Информация об оплате
	Items             []Item    `json:"items"`              // Список элементов заказа
	Locale            string    `json:"locale"`             // Локаль заказа
	InternalSignature string    `json:"internal_signature"` // Внутренняя подпись
	CustomerId        string    `json:"customer_id"`        // Идентификатор клиента
	DeliveryService   string    `json:"delivery_service"`   // Служба доставки
	Shardkey          string    `json:"shardkey"`           // Ключ шарда
	SmId              int       `json:"sm_id"`              // Идентификатор службы маркетинга
	DateCreated       time.Time `json:"date_created"`       // Дата и время создания заказа
	OofShard          string    `json:"oof_shard"`          // OOF Shard
}
