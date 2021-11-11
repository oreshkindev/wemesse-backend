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

// опрееляем модель, которую будем возвращать на клиент
func (dbUser *Users) UsersToWeb() *Users {
	return &Users{
		ID:              dbUser.ID,
		AppLocale:       dbUser.AppLocale,
		AppVersion:      dbUser.AppVersion,
		DeviceLocale:    dbUser.DeviceLocale,
		DeviceMac:       dbUser.DeviceMac,
		DeviceModel:     dbUser.DeviceModel,
		DeviceSDK:       dbUser.DeviceSDK,
		SessionActivity: dbUser.SessionActivity,
		SessionID:       dbUser.SessionID,
		SessionRegister: dbUser.SessionRegister,
		TgVersion:       dbUser.TgVersion,
	}
}
