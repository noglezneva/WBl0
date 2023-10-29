package adapter

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/noglezneva/wbl0/internal/model"
)

const path = "/build/view/"

// CacheService определяет методы для работы с кэшем.
type CacheService interface {
	Set(int64, *model.Order) bool
	SetAll(orders []*model.Order)
	Get(int64) *model.Order
}

// DbService определяет методы для работы с базой данных
type DbService interface {
	GetByOrderUID(string) (*model.Order, error)
	GetByID(int64) (*model.Order, error)
	Create(*model.Order) error
	GetAll() ([]*model.Order, error)
}

// Adapter представляет собой адаптер, который обеспечивает взаимодействие между кэшем и базой данных.
type Adapter struct {
	cache CacheService
	db    DbService
}

// NewAdapter создает новый экземпляр Adapter.
func NewAdapter(cache CacheService, db DbService) *Adapter {
	return &Adapter{cache: cache, db: db}
}

// SetAllCacheFromDB устанавливает все заказы в кэш из базы данных.
func (a *Adapter) SetAllCacheFromDB() error {
	orders, err := a.db.GetAll()
	if err != nil {
		log.Println(err)
		return err
	}
	a.cache.SetAll(orders)
	return nil
}

// HandleGet обрабатывает GET-запрос. Извлекает id заказа из параметров запроса,
// пытается найти его в кэше, если не найден - обращается к базе данных.
func (a *Adapter) HandleGet(w http.ResponseWriter, r *http.Request) {
	var id int64
	if err := extractParamInt64(&id, "id", r); err != nil {
		drawError(w)
		return
	}
	log.Println("get id ", id)

	var order *model.Order
	order = a.cache.Get(id)
	if order != nil {
		drawOrder(w, order)
	}

	order, err := a.db.GetByID(id)
	if err != nil {
		log.Println("db error")
		log.Println(err)
		drawError(w)
		return
	}

	go a.cache.Set(id, order)

	drawOrder(w, order)
}

// extractParamInt64 извлекает целочисленное значение из параметра запроса и сохраняет его в переменной i.
func extractParamInt64(i *int64, paramName string, r *http.Request) error {
	paramString := r.FormValue(paramName)
	param, err := strconv.ParseInt(paramString, 10, 32)
	if err != nil {
		return err
	}
	*i = param
	return nil
}

// drawError отображает страницу с сообщением об ошибке.
func drawError(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(path + "getError.html")
	if err != nil {
		log.Println("tmpl parse error")
		log.Println(err)
		return
	}
	err = tmpl.Execute(w, "")
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// drawOrder генерирует HTML-представление для заказа и отображает его на странице.
func drawOrder(w http.ResponseWriter, order *model.Order) {
	tmpl, err := template.ParseFiles(path + "get.html")
	if err != nil {
		log.Println("tmpl parse error")
		log.Println(err)
		return
	}
	err = tmpl.Execute(w, order)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
