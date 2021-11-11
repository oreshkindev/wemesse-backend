// Package main provides
package main

import (
	"context"
	"log"
	"net/http"
	"wemesse/conf"
	"wemesse/controller"
	"wemesse/router"
	"wemesse/service"
	"wemesse/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// инициализируем приложение
func run() error {
	ctx := context.Background()

	// инициализируем наш конфиг
	conf, err := conf.New()
	if err != nil {
		return err
	}

	// инициализируем наш репозиторий
	store, err := store.New(ctx, &conf.Postgres)
	if err != nil {
		return err
	}

	// инициализируем наши сервисы
	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		return err
	}

	// инициализируем наши контроллеры
	updatesController := controller.NewUpdates(ctx, serviceManager, conf)
	usersController := controller.NewUser(ctx, serviceManager)

	// инициализируем роутер
	router, err := router.NewRouter(ctx)
	if err != nil {
		return err
	}

	// инициализируем маршруты
	// получаем версию для старых пользователей
	router.Get("/api/v1/updates/{version}", func(w http.ResponseWriter, r *http.Request) {
		updatesController.GetUpdate(w, r)
	})

	// служебный метод
	// получаем списки обновлений и статистику скачиваний
	router.Get("/api/v1/updates", func(w http.ResponseWriter, r *http.Request) {
		updatesController.GetUpdates(w, r)
	})

	// служебный метод
	// пушим приложение на сервер
	router.Post("/api/v1/updates", func(w http.ResponseWriter, r *http.Request) {
		updatesController.PostUpdate(w, r)
	})

	// добавляем нового пользователя + проверки и обновление
	router.Post("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		usersController.PostUser(w, r)
	})

	// служебный метод
	// получаем списки обновлений и статистику скачиваний
	router.Get("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		usersController.GetUsers(w, r)
	})

	// запускаем приложение если все ok
	if err := http.ListenAndServe(":"+conf.Port, router); err != nil {
		return err
	}

	return nil
}
