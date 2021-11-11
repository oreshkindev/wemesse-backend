// Package model provides
package model

// определяем структуру нашего хранилища *(таблицы) и используем теги для структур для переопределения имен полей
type Middleware struct {
	ID        int
	Signature string // Argon2 Hash
}
