package model

// Delivery представляет информацию о доставке.
type Delivery struct {
	ID      int64  // Уникальный идентификатор доставки
	Name    string `json:"name"`    // Имя получателя
	Phone   string `json:"phone"`   // Телефон получателя
	Zip     string `json:"zip"`     // Почтовый индекс
	City    string `json:"city"`    // Город
	Address string `json:"address"` // Адрес доставки
	Region  string `json:"region"`  // Регион
	Email   string `json:"email"`   // Email получателя
}
