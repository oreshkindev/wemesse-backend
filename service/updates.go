// Package service provides
package service

import (
	"context"
	"wemesse/model"
	"wemesse/store"
)

// определяем структуру
type UpdatesWebService struct {
	ctx   context.Context
	store *store.Store
}

// создаем наш сервис проверки обновлений
func NewUpdatesWebService(ctx context.Context, store *store.Store) *UpdatesWebService {
	return &UpdatesWebService{
		ctx:   ctx,
		store: store,
	}
}

// опрабатываем запрос от контроллера через интерфейс
func (svc *UpdatesWebService) GetUpdate(ctx context.Context, appVersion string) (*model.App, error) {

	// передаем полученную версию в интерфейс хранилища
	updatesDB, err := svc.store.Updates.GetUpdate(ctx, appVersion)
	if err != nil {
		return nil, err
	}
	if updatesDB == nil {
		return nil, err
	}

	// вернем в контроллер наш результат
	return updatesDB, nil
}

// служебный метод
// опрабатываем запрос от контроллера через интерфейс
func (svc *UpdatesWebService) GetUpdates(ctx context.Context) ([]model.App, error) {

	// передаем полученную версию в интерфейс хранилища
	updatesDB, err := svc.store.Updates.GetUpdates(ctx)
	if err != nil {
		return nil, err
	}
	if updatesDB == nil {
		return nil, err
	}

	// вернем в контроллер наш результат
	return updatesDB, nil
}

// служебный метод
// опрабатываем запрос от контроллера через интерфейс
func (svc *UpdatesWebService) PostUpdate(ctx context.Context, app *model.App) (*model.App, error) {

	// передаем полученную версию в интерфейс хранилища
	updatesDB, err := svc.store.Updates.PostUpdate(ctx, app)
	if err != nil {
		return nil, err
	}
	if updatesDB == nil {
		return nil, err
	}

	// вернем в контроллер наш результат
	return updatesDB, nil
}

// служебный метод
// опрабатываем запрос от контроллера через интерфейс
func (svc *UpdatesWebService) VerifySignature(ctx context.Context, signature string) error {

	// передаем полученную версию в интерфейс хранилища
	err := svc.store.Updates.GetSignature(ctx, signature)
	if err != nil {
		return err
	}

	// вернем в контроллер наш результат
	return nil
}
