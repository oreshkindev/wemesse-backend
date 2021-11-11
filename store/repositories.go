// Package store provides
package store

import (
	"context"
	"wemesse/model"
)

// UpdatesRepo - интерфейс к которому обращается наш сервис
type UpdatesRepo interface {
	GetUpdate(context.Context, string) (*model.App, error)
	GetUpdates(context.Context) ([]model.App, error)
	GetSignature(context.Context, string) error
	PostUpdate(context.Context, *model.App) (*model.App, error)
}

// UsersRepo - интерфейс к которому обращается наш сервис
type UsersRepo interface {
	GetUsers(context.Context) ([]model.Users, error)
	PostUser(context.Context, *model.Users) (*model.Users, error)
}
