package main

import (
	"log"
	"net/http"

	"github.com/noglezneva/wbl0/internal/adapter"
	"github.com/noglezneva/wbl0/internal/model"
	nats "github.com/noglezneva/wbl0/internal/nats-streaming"
	"github.com/noglezneva/wbl0/internal/rest"
	cache2 "github.com/noglezneva/wbl0/internal/service/cache"
	db2 "github.com/noglezneva/wbl0/internal/service/db"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// объявляем константы
const port = "8080"
const cacheCapacity = 10000

func main() {
	//устанавливаем соединение с базой данных
	conn := pg.Connect(&pg.Options{
		User:     "root",
		Password: "5432",
		Database: "postgres",
		Addr:     "db:5432",
		//Addr: "localhost:5433",
	})

	err := createSchema(conn)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	cache := cache2.NewLru(cacheCapacity)
	db := db2.NewService(conn)
	ad := adapter.NewAdapter(cache, db)

	err = ad.SetAllCacheFromDB()
	if err != nil {
		log.Println(err)
	}

	nats.SubscribeAndPublishTest(db)

	//создаем маршрутизатор для обработки HTTP-запросов
	router := rest.Handler(ad)
	log.Printf("Starting server on port %s.....\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.Delivery)(nil),
		(*model.Payment)(nil),
		(*model.Item)(nil),
		(*model.Order)(nil),
	}

	for _, m := range models {
		// создаем таблицы в бд
		err := db.Model(m).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
