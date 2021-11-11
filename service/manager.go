// Package service provides
package service

import (
	"context"
	"errors"
	"wemesse/store"
)

// определим структуру со всеми сервисами с которыми мы будем взаимодействовать
type Manager struct {
	Updates UpdatesService
	Users   UsersService
}

// создадим сервис менеджер
func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("NO STORE PROVIDED")
	}
	// заполним структуру менеджера
	return &Manager{
		Updates: NewUpdatesWebService(ctx, store),
		Users:   NewUsersWebService(ctx, store),
	}, nil
}
