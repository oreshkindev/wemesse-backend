// Package postgres provides
package postgres

import (
	"context"
	"strings"
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

	var targetABI string = ""

	if len(appVersion) == 10 || len(appVersion) == 12 {

		if len(appVersion) == 10 && strings.Contains(appVersion[len(appVersion)-3:], "arm") {
			targetABI = appVersion[len(appVersion)-3:]
		}
		if len(appVersion) == 12 && strings.Contains(appVersion[len(appVersion)-5:], "arm64") {
			targetABI = appVersion[len(appVersion)-5:]
		}

	}
	// определяем структуру для объекта
	tmp := &model.App{}

	prepare := &model.App{}

	// достаем последнюю запись из хранилища где skipped = false
	if r := repo.db.Last(&tmp, "skipped = ?", false); r.Error != nil {
		return nil, r.Error
	}

	if targetABI != "" {
		// достаем последнюю запись архитектуры билда из хранилища
		if r := repo.db.Last(&prepare, "target_abi", targetABI); r.Error != nil {
			return nil, r.Error
		}
	} else {
		// достаем последнюю запись из хранилища
		if r := repo.db.Last(&prepare, "target_abi", "arm64"); r.Error != nil {
			return nil, r.Error
		}
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
// сигнатура необходима для верификации клиента
// достаем последнюю запись с сигнатурой из хранилища
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
	if repo.db.Model(&app).Where("app_version = ? AND target_abi = ?", app.AppVersion, app.TargetABI).Updates(&app).RowsAffected == 0 {
		repo.db.Create(&app)
	}

	// возвращаем наше обновление
	return app, nil
}
