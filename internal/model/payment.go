package model

// Payment представляет информацию об оплате.
type Payment struct {
	ID           int64  // Уникальный идентификатор оплаты
	Transaction  string `json:"transaction"`   // Идентификатор транзакции
	RequestId    string `json:"request_id"`    // Идентификатор запроса
	Currency     string `json:"currency"`      // Валюта оплаты
	Provider     string `json:"provider"`      // Провайдер оплаты
	Amount       int    `json:"amount"`        // Сумма оплаты
	PaymentDt    int    `json:"payment_dt"`    // Дата и время оплаты
	Bank         string `json:"bank"`          // Банк
	DeliveryCost int    `json:"delivery_cost"` // Стоимость доставки
	GoodsTotal   int    `json:"goods_total"`   // Общая стоимость товаров
	CustomFee    int    `json:"custom_fee"`    // Плата за обработку заказа
}
