// Package postgres provides
package postgres

import (
	"context"
	"wemesse/model"
)

// определяем структуру
type UsersRepo struct {
	db *DB
}

// создаем наш непозиторий с пользователями
func NewUsersRepo(db *DB) *UsersRepo {
	return &UsersRepo{db: db}
}

// служебный метод
// достаем списки пользователей из хранилища
func (repo *UsersRepo) GetUsers(ctx context.Context) ([]model.Users, error) {
	// определяем структуру массива
	arr := []model.Users{}

	// достаем все записи из таблицы apps
	if r := repo.db.Raw("select * from users order by id asc").Scan(&arr); r.Error != nil {
		return nil, r.Error
	}

	// возвращаем наши записи
	return arr, nil
}

// TODO:
// убрать взаимодействие с таблицей apps
// и вынести count +1 в отдельный сервис
// или контроллер

// получем данные от сервиса через интерфейс хранилища и пишем в базу нашего пользователя
func (repo *UsersRepo) PostUser(ctx context.Context, user *model.Users) (*model.Users, error) {

	var prepare model.Users
	// пытаемся найти юзера по id сессии и мак адресу
	r := repo.db.First(&prepare, "session_id = ? and device_mac = ?", user.SessionID, user.DeviceMac)
	// если не находим совпадение, то ...
	if r.Error != nil {
		// создаем запись в хранилище
		if r := repo.db.Create(&user); r.Error != nil {
			return nil, r.Error
		}
		// обновляем кол-во скачиваний для нового пользователя или нового устройства
		if r := repo.db.Raw("update apps set uploads = uploads + 1 where app_version = ? AND target_abi = ?", user.AppVersion, user.TargetABI).Scan(&user); r.Error != nil {
			return nil, r.Error
		}
	} else {
		// обновляем запись в хранилище по мак адресу устройства
		if r := repo.db.Raw("update users set app_locale = ?, app_version = ?, device_locale = ?, device_sdk = ?, session_activity = ?, session_register = ?, tg_version = ?, target_abi = ? where device_mac = ? RETURNING id", user.AppLocale, user.AppVersion, user.DeviceLocale, user.DeviceSDK, user.SessionActivity, user.SessionRegister, user.TgVersion, user.TargetABI, user.DeviceMac).Scan(&user); r.Error != nil {
			return nil, r.Error
		}
		// проверяем небыло ли у пользователя такой версии приложения ранее, и если нет, то ...
		if user.AppVersion != prepare.AppVersion {
			// обновляем кол-во скачиваний для нового пользователя или нового устройства
			if r := repo.db.Raw("update apps set uploads = uploads + 1 where app_version = ? AND target_abi = ? RETURNING id", user.AppVersion, user.TargetABI).Scan(&user); r.Error != nil {
				return nil, r.Error
			}
		}
	}

	// возвращаем нашего пользователя
	return user, nil
}
