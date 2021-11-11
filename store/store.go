// Package store provides
package store

import (
	"context"

	"wemesse/conf"
	"wemesse/store/postgres"
)

// определим структуру со всеми репозиториями с которыми мы будем взаимодействовать
type Store struct {
	Pg *postgres.DB // for KeepAliveGorm (see below)

	Updates UpdatesRepo
	Users   UsersRepo
}

// создаем новое хранилище, сервисы и миграцию
func New(ctx context.Context, conf *conf.Postgres) (*Store, error) {

	// коннект к хранилищу
	conn, err := postgres.Dial(conf)
	if err != nil {
		return nil, err
	}

	// запускаем миграцию
	if conn != nil {
		if err := migrate(conn); err != nil {
			return nil, err
		}
	}

	// определяем хранилище и его структуру
	var store Store

	// заполняем структуру
	if conn != nil {
		store.Pg = conn
		store.Updates = postgres.NewUpdatesRepo(conn)
		store.Users = postgres.NewUsersRepo(conn)
	}

	// возвращаем хранилище
	return &store, nil
}
