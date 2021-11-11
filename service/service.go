// Package service provides
package service

import (
	"context"
	"wemesse/model"
)

// UpdatesService - интерфейс к которому обращается наш котроллер
type UpdatesService interface {
	GetUpdate(context.Context, string) (*model.App, error)
	GetUpdates(context.Context) ([]model.App, error)
	PostUpdate(context.Context, *model.App) (*model.App, error)
	VerifySignature(context.Context, string) error
}

// UsersService - интерфейс к которому обращается наш котроллер
type UsersService interface {
	GetUsers(context.Context) ([]model.Users, error)
	PostUser(context.Context, *model.Users) (*model.Users, error)
}
