// Package conf provides
package conf

import "os"

// определяем структуру конфигурации приложения
type Conf struct {
	Host      string
	Port      string
	DestURI   string
	SourceURI string
	Postgres  Postgres
}

// Postgres - определяем структуру конфигурации хранилища
type Postgres struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

// создаем новый конфиг и заполняем нашу структуру данными из env.sh
func New() (*Conf, error) {
	return &Conf{
		// "переменная", "дефолтное значение"
		Host:      env("app_host", "127.0.0.1"),
		Port:      env("app_port", "8080"),
		DestURI:   env("dest_uri", "http://182.92.107.179/wemesse/source/"),
		SourceURI: env("source_uri", "https://messenger.tbcc.com/source/"),
		Postgres: Postgres{
			User: env("user", "postgres"),
			Pass: env("pass", "postgres"),
			Host: env("host", "127.0.0.1"),
			Port: env("port", "5432"),
			Name: env("name", "postgres"),
		},
	}, nil
}

// собираем переменные окружения из env.sh
func env(key, def string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return v
}
