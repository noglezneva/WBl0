package rest

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/noglezneva/wbl0/internal/adapter"

	"github.com/go-chi/chi"
)

const path = "/build/view/"

// обработчик запросов
func Handler(a *adapter.Adapter) *chi.Mux {
	router := chi.NewRouter()

	// Определение маршрутов и обработчиков для главной страницы и запроса GET
	router.Route("/", func(r chi.Router) {
		// Обработчик для главной страницы
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			HandlerIndex(w)
		})
		// Обработчик для запроса GET
		r.Get("/get", a.HandleGet)
	})
	return router
}

// Обработчик главной страницы
func HandlerIndex(out io.Writer) {
	// Загружаем шаблон HTML страницы
	tmpl, err := template.ParseFiles(path + "index.html")
	if err != nil {
		log.Println("tmpl parse error")
		log.Println(err)
		return
	}
	// Отправляем сгенерированную страницу в ответ
	err = tmpl.Execute(out, "")
	if err != nil {
		return
	}
}
