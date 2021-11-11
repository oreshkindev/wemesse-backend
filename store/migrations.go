// Package postgres provides
package store

import (
	"wemesse/model"
	"wemesse/store/postgres"
)

// миграция моделей в хранилище
func migrate(db *postgres.DB) error {
	// передаем необходимые модели для миграции
	if err := db.AutoMigrate(&model.App{}, &model.Users{}, &model.Middleware{}); err != nil {
		return err
	}
	// ничего не возвращаем т.к нет смысла
	return nil
}
