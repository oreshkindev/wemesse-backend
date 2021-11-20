// Package model provides
package model

// определяем структуру нашего хранилища *(таблицы) и используем теги для структур для переопределения имен полей
type App struct {
	ID          int    `json:"Id"`
	AppChecksum string `json:"Checksum"`
	AppName     string `json:"Name"`
	AppNotes    string `json:"Message"`
	AppSize     string `json:"Size"`
	AppVersion  string `json:"Version"`
	Skipped     bool   `json:"Skipped"`
	Uploads     int    `json:"Uploads"`
	URI         string `json:"URI"`
}
