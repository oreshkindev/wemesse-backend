// Package postgres provides
package postgres

import (
	"context"
	"wemesse/model"
)

// определяем структуру
type UpdatesRepo struct {
	db *DB
}

// создаем наш репозиторий с пользователями
func NewUpdatesRepo(db *DB) *UpdatesRepo {
	return &UpdatesRepo{db: db}
}

// достаем последнюю запись с обновлением из хранилища
func (repo *UpdatesRepo) GetUpdate(ctx context.Context, appVersion string) (*model.App, error) {
	// определяем структуру для объекта
	tmp := &model.App{}

	prepare := &model.App{}

	// достаем последнюю запись из хранилища где skipped = false
	if r := repo.db.Last(&tmp, "skipped = ?", false); r.Error != nil {
		return nil, r.Error
	}

	// достаем последнюю запись из хранилища
	if r := repo.db.Last(&prepare); r.Error != nil {
		return nil, r.Error
	}

	if tmp.AppVersion >= appVersion {

		prepare.Skipped = false
	}

	// возвращаем наше последнее обновление
	return prepare, nil
}

// служебный метод
// достаем списки обновлений из хранилища
func (repo *UpdatesRepo) GetUpdates(ctx context.Context) ([]model.App, error) {
	// определяем структуру массива
	arr := []model.App{}

	// достаем все записи из таблицы apps
	if r := repo.db.Raw("select * from apps order by id asc").Scan(&arr); r.Error != nil {
		return nil, r.Error
	}

	// возвращаем наши записи
	return arr, nil
}

// служебный метод
// достаем последнюю запись с обновлением из хранилища
func (repo *UpdatesRepo) GetSignature(ctx context.Context, signature string) error {

	var key model.Middleware

	// пытаемся найти юзера по id сессии и мак адресу
	r := repo.db.First(&key, "signature = ?", signature)
	// если не находим совпадение, то ...
	if r.Error != nil {
		// создаем запись в хранилище
		return r.Error
	}

	// возвращаем наше последнее обновление
	return nil
}

// служебный метод
// получем данные от сервиса через интерфейс хранилища и пишем в базу нашего пользователя
func (repo *UpdatesRepo) PostUpdate(ctx context.Context, app *model.App) (*model.App, error) {

	// пишем в хранилище обновление
	if r := repo.db.Create(&app); r.Error != nil {
		return nil, r.Error
	}

	// возвращаем нашего пользователя
	return app, nil
}
