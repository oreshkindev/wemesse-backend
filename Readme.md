Логика:

```bash
Таблица app: {
"appName"
"appSize"
"appVersion"
"checksum"
"notes"
"uploads"
"URI" - возвращаем директорию бэкенда при посте для Сереги, а в приложение уже возвращаем ссылку на Серегу.
}
```

```bash
Таблица users: {
"ID"
"appLocale"
"appVersion" - Если записаная версия == той что получили, ничего не делаем, иначе count + 1 к таблице app -> uploads новой версии юзера
"deviceLocale"
"deviceMac"
"deviceModel"
"deviceSDK"
"sessionActivity"
"sessionID"
"sessionRegister"
"tgVersion"
}
```

\*\*\* Если на бэк прилетает GET мы перманентно отправляем ему "skipped": false, чтобы он полюбому обновился.

```bash
- POST Получаем JSON с клиента {
  "appLocale"
  "appVersion"
  "deviceLocale"
  "deviceMac"
  "deviceModel"
  "deviceSDK"
  "sessionActivity"
  "sessionID"
  "sessionRegister"
  "tgVersion"
  }
```

```bash
- Проверяем appVersion (проверяем строку из POST запроса) -> {
  "appName": "wemese.apk",
  "appSize": "0 B",
  "appVersion": "2264_00",
  "checksum": "d41d8cd98f00b204e9800998ecf8427e",
  "notes": "",
  "skipped": false,
  "URI": "https://seregin.server/release/wemese.apk"
  }
```

```bash
<!-- GET https://messenger.tbcc.com/api/v1/updates/2264_00 -> {
  "appName": "wemese.apk",
  "appSize": "0 B",
  "appVersion": "2264_00",
  "checksum": "d41d8cd98f00b204e9800998ecf8427e",
  "notes": "",
  "skipped": false,
  "URI": "https://seregin.server/release/wemese.apk"
  } -->
 ```
