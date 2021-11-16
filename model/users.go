// Package model provides
package model

// определяем структуру нашего хранилища *(таблицы)
type Users struct {
	ID              int
	AppLocale       string
	AppVersion      string
	DeviceLocale    string
	DeviceMac       string
	DeviceModel     string
	DeviceSDK       string
	SessionActivity int
	SessionID       int
	SessionRegister int
	TgVersion       string
}
