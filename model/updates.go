// Package model provides
package model

// определяем структуру нашего хранилища *(таблицы) и используем теги для структур для переопределения имен полей
type App struct {
	ID          int    `json:"Id"`
	AppChecksum string `json:"Checksum"`
	AppName     string `json:"Name"`
	AppNotes    string `json:"Message"`
	AppSize     int64  `json:"Size"`
	AppVersion  string `json:"Version"`
	Skipped     bool   `json:"Skipped"`
	Uploads     int    `json:"Uploads"` // игнорируем и не возвращаем поле
	URI         string `json:"URI"`
}

// опрееляем модель, которую будем возвращать на клиент
func (dbUser *App) ToWeb() *App {
	return &App{
		ID:          dbUser.ID,
		AppChecksum: dbUser.AppChecksum,
		AppNotes:    dbUser.AppNotes,
		AppName:     dbUser.AppName,
		AppSize:     dbUser.AppSize,
		AppVersion:  dbUser.AppVersion,
		Skipped:     dbUser.Skipped,
		URI:         dbUser.URI,
	}
}
