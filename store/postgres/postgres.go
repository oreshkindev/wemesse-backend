// Package postgres provides
package postgres

import (
	"fmt"
	"wemesse/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// определяем структуру
type DB struct {
	*gorm.DB
}

// конектимся к хранилищу
func Dial(c *conf.Postgres) (*DB, error) {

	// подставляем полученные параметры в строку
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Host, c.Port, c.User, c.Name, c.Pass)

	// создаем подключение
	conn, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		return nil, err
	}

	// возвращаем подключение
	return &DB{conn}, nil
}
