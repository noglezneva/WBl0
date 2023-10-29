package model

// Item представляет информацию об элементе.
type Item struct {
	ID          int64  // Уникальный идентификатор элемента
	ChrtId      int    `json:"chrt_id"`      // Идентификатор корзины
	TrackNumber string `json:"track_number"` // Номер отслеживания
	Price       int    `json:"price"`        // Цена
	Rid         string `json:"rid"`          // Идентификатор объекта
	Name        string `json:"name"`         // Название элемента
	Sale        int    `json:"sale"`         // Скидка
	Size        string `json:"size"`         // Размер
	TotalPrice  int    `json:"total_price"`  // Общая стоимость
	NmId        int    `json:"nm_id"`        // Идентификатор торговой марки
	Brand       string `json:"brand"`        // Бренд
	Status      int    `json:"status"`       // Статус
}
