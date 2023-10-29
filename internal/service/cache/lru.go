package cache

import (
	"container/list"
	"sync"

	"github.com/noglezneva/wbl0/internal/model"
)

// Item представляет элемент кэша, состоящий из ключа типа int64 и значения типа *model.Order
type Item struct {
	Key   int64
	Value *model.Order
}

// LRU представляет реализацию алгоритма LRU кэша.
type LRU struct {
	capacity int                     // емкость кэша
	items    map[int64]*list.Element // хэш-таблица соответствия ключа и элемента в очереди
	queue    *list.List              // двусвязный список для отслеживания порядка использования элементов
	mu       sync.RWMutex            // мьютекс для обеспечения потокобезопасности доступа к кэшу
}

// NewLru создает новый экземпляр LRU кэша с указанной емкостью.
func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[int64]*list.Element),
		queue:    list.New(),
		mu:       sync.RWMutex{},
	}
}

// Set добавляет элемент в кэш или обновляет его значение.
func (c *LRU) Set(key int64, value *model.Order) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if element, exists := c.items[key]; exists {
		// Если элемент уже существует, перемещаем его в начало очереди
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	if c.queue.Len() == c.capacity {
		// Если емкость кэша достигнута, удаляем самый старый элемент
		c.purge()
	}

	// Создаем новый элемент и добавляем его в начало очереди
	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element

	return true
}

// SetAll добавляет все заказы из переданного списка в кэш
func (c *LRU) SetAll(orders []*model.Order) {
	for _, value := range orders {
		c.Set(value.ID, value)
	}
}

// purge удаляет самый старый элемент из кэша
func (c *LRU) purge() {
	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*Item)
		delete(c.items, item.Key)
	}
}

// Get возвращает значение, связанное с указанным ключом в кэше
func (c *LRU) Get(key int64) *model.Order {
	c.mu.RLock()
	defer c.mu.RUnlock()
	element, exists := c.items[key]
	if !exists {
		return nil
	}
	return element.Value.(*Item).Value
}
