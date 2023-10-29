package db

import (
	"log"

	"github.com/noglezneva/wbl0/internal/model"

	"github.com/go-pg/pg/v10"
)

// Service представляет сервис базы данных
type Service struct {
	db *pg.DB // Клиент базы данных
}

// NewService создает новый экземпляр сервиса базы данных с указанным клиентом базы данных
func NewService(db *pg.DB) *Service {
	return &Service{db: db}
}

// GetByOrderUID получает заказ из базы данных по его идентификатору OrderUID
func (s *Service) GetByOrderUID(orderUID string) (*model.Order, error) {
	order := &model.Order{OrderUID: orderUID}
	err := s.db.Model(order).Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return order, nil
}

// GetByID получает заказ из базы данных по его идентификатору ID
func (s *Service) GetByID(id int64) (*model.Order, error) {
	order := &model.Order{ID: id}
	err := s.db.Model(order).WherePK().Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return order, nil
}

// GetAll получает все заказы из базы данных.
func (s *Service) GetAll() ([]*model.Order, error) {
	var order []*model.Order
	err := s.db.Model(&order).Select()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return order, nil
}

// Create создает новый заказ в базе данных
func (s *Service) Create(order *model.Order) error {
	_, err := s.db.Model(order).Insert()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
