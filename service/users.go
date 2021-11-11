// Package service provides
package service

import (
	"context"
	"wemesse/model"
	"wemesse/store"
)

// определяем структуру
type UsersWebService struct {
	ctx   context.Context
	store *store.Store
}

// создаем наш сервис для пользователей
func NewUsersWebService(ctx context.Context, store *store.Store) *UsersWebService {
	return &UsersWebService{
		ctx:   ctx,
		store: store,
	}
}

// служебный метод
// опрабатываем запрос от контроллера через интерфейс
func (svc *UsersWebService) GetUsers(ctx context.Context) ([]model.Users, error) {

	// передаем полученную версию в интерфейс хранилища
	usersDB, err := svc.store.Users.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	if usersDB == nil {
		return nil, err
	}

	// вернем в контроллер наш результат
	return usersDB, nil
}

// получаем тело запроса от контроллера через интерфейс
func (svc *UsersWebService) PostUser(ctx context.Context, user *model.Users) (*model.Users, error) {
	// передаем полученную модель в интерфейс хранилища
	updatesDB, err := svc.store.Users.PostUser(ctx, user)
	if err != nil {
		return nil, err
	}
	if updatesDB == nil {
		return nil, err
	}

	// вернем в контроллер наш результат
	return updatesDB.UsersToWeb(), nil
}
