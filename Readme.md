## Хранилище

После первого запуска бинарника происходит миграция всех импортированных моделей в хранилище

### Таблица apps:

Основное хранилище обновлений и версионирования клиента

```bash
{
    "id":               bigint,
    "app_checksum":     text,
    "app_name":         text,
    "app_notes":        text,
    "app_size":         bigint,
    "app_version":      text,
    "skipped":          boolean,
    "uploads":          bigint,
    "uri":              text
}
```

### Таблица middlewares:

Необходима для сверки сигнатуры клиента и сервера. Сигнатура создается принудительно в хранилище используя Argon2 или любой другой алгоритм

```bash
{
    "id":               bigint,
    "signature":        text
}
```

### Таблица users:

Основное хранилище пользователей передаваемых при каждой сессии с клиента

```bash
{
    "id":                   bigint,
    "app_locale":           text,
    "app_version":          text,
    "device_locale":        text,
    "device_mac":           text,
    "device_model":         text,
    "device_sdk":           text,
    "session_activity":     bigint,
    "session_id":           bigint,
    "session_register":     bigint,
    "tg_version":           text
}
```

## Методы

Основные методы. Все за исключением GetUpdate ожидают сигнатуру клиента

### GET / GetUpdate

Передаем текущую версию клиента. Возвращаем JSON с актуальной версией или принудительной для загрузки

```bash
router.Get("/api/v1/updates/{version}"
```

### Response:

ВАЖНО: URI вернется с ссылкой Гонконгского сервера

```bash
{
    "Id":           int,
    "Checksum":     string,
    "Name":         string,
    "Message":      string,
    "Size":         int64,
    "Version":      string,
    "Skipped":      bool,
    "Uploads":      int,
    "URI":          string
}
```

### GET / GetUpdates

ВАЖНО: Служебный метод
Возвращаем массив объектов со всеми обновлениями

```bash
router.Get("/api/v1/updates"
```

### Response:

ВАЖНО: URI вернется с ссылкой нашего сервера

```bash
[
    {
        "Id":           int,
        "Checksum":     string,
        "Name":         string,
        "Message":      string,
        "Size":         int64,
        "Version":      string,
        "Skipped":      bool,
        "Uploads":      int,
        "URI":          string
    }
]
```

### POST / PostUpdate

ВАЖНО: Служебный метод
ВАЖНО: Каждая 00_00 версия клиента принудительная и вернется со статусом skipped = false
Передаем Form-Data. Возвращаем JSON этого обновления

```bash
router.Post("/api/v1/updates"

appNotes        string
appVersion      string
tmpFile         file
```

### Response:

ВАЖНО: URI вернется с ссылкой нашего сервера

```bash
{
    "Id":           int,
    "Checksum":     string,
    "Name":         string,
    "Message":      string,
    "Size":         int64,
    "Version":      string,
    "Skipped":      bool,
    "Uploads":      int,
    "URI":          string
}
```

### POST / PostUser

ВАЖНО: Этот метод прибавляет +1 к версии в таблице apps
Передаем с клиента JSON объект пользователя и его устройства. Возвращаем JSON объект текущего пользователя

```bash
router.Post("/api/v1/users"

{
    "AppLocale":            string,
    "AppVersion":           string,
    "DeviceLocale":         string,
    "DeviceMac":            string,
    "DeviceModel":          string,
    "DeviceSDK":            string,
    "SessionActivity":      int,
    "SessionID":            int,
    "SessionRegister":      int,
    "TgVersion":            string
}
```

### Response:

```bash
{
    "ID":                   int,
    "AppLocale":            string,
    "AppVersion":           string,
    "DeviceLocale":         string,
    "DeviceMac":            string,
    "DeviceModel":          string,
    "DeviceSDK":            string,
    "SessionActivity":      int,
    "SessionID":            int,
    "SessionRegister":      int,
    "TgVersion":            string
}
```

### GET / GetUsers

ВАЖНО: Служебный метод
Возвращаем массив объектов со всеми пользователями

```bash
router.Get("/api/v1/users"
```

### Response:

```bash
[
    {
        "ID":                   int,
        "AppLocale":            string,
        "AppVersion":           string,
        "DeviceLocale":         string,
        "DeviceMac":            string,
        "DeviceModel":          string,
        "DeviceSDK":            string,
        "SessionActivity":      int,
        "SessionID":            int,
        "SessionRegister":      int,
        "TgVersion":            string
    }
]
```
