# loglint

Линтер для Go, проверяет лог-сообщения на соответствие правилам стиля и безопасности.
Совместим с golangci-lint в качестве плагина.

## Правила

| Правило | Пример ❌ | Пример ✅ |
|---|---|---|
| Сообщение начинается со строчной буквы | `"Starting server"` | `"starting server"` |
| Только английский язык | `"запуск сервера"` | `"starting server"` |
| Нет спецсимволов и эмодзи | `"server started!🚀"` | `"server started"` |
| Нет чувствительных данных | `"password: " + pass` | `"user authenticated"` |

Поддерживаемые логгеры: `log/slog`, `go.uber.org/zap`

## Структура проекта
```
loglint/
├── cmd/loglint/
│   └── main.go         # точка входа для standalone-запуска
├── internal/
│   ├── analyzer/
│   │   └── analyzer.go # обход AST, вызов правил
│   ├── rules/
│   │   └── rules.go    # логика проверок
│   └── utils/
│       └── utils.go    # хелперы: определение лог-вызова, извлечение сообщения
├── testdata/src/
│   ├── slog/
│   │   └── slog.go     # тестовые случаи для log/slog
│   └── uberzap/
│       └── zap.go      # тестовые случаи для go.uber.org/zap
├── .golangci.yml
├── Makefile
└── README.md
```

## Требования

- Go 1.22+
- golangci-lint (если нужна интеграция)

## Установка
```bash
git clone https://github.com/s0mewha7/loglint
cd loglint
go mod tidy
```

## Сборка и запуск

### Standalone
```bash
make build
./bin/loglint ./...
```

Или без сборки:
```bash
go run ./cmd/loglint ./...
```

### Плагин для golangci-lint

Собрать `.so`:
```bash
make plugin
```

Добавить в `.golangci.yml` своего проекта:
```yaml
linters-settings:
  custom:
    loglint:
      path: /path/to/loglint/bin/loglint.so
      description: "checks log messages for style and safety issues"

linters:
  enable:
    - loglint
```

Запустить:
```bash
golangci-lint run ./...
```

## Тесты
```bash
make test
```

## Примеры
```go
// ❌ не пройдёт проверку

slog.Info("Starting server on port 8080")  // заглавная буква
zap.Error("Failed to connect")             // заглавная буква

slog.Info("запуск сервера")                // не английский язык
zap.Error("ошибка подключения")            // не английский язык

slog.Info("server started!🚀")             // эмодзи и спецсимволы
zap.Error("connection failed!!!")           // спецсимволы

slog.Info("user password: " + password)    // чувствительные данные
zap.Debug("api_key=" + apiKey)             // чувствительные данные

// ✅ пройдёт проверку

slog.Info("starting server on port 8080")
zap.Error("failed to connect to database")
slog.Info("user authenticated successfully")
zap.Debug("api request completed")
```

Вывод линтера:
```
./main.go:6:2: log message must start with lowercase
./main.go:9:2: log message must be in english
./main.go:12:2: log message must not contain special characters
./main.go:15:2: possible sensitive data in logs: password
```

## Makefile

| Команда | Описание |
|---|---|
| `make` / `make build` | Собрать бинарник в `bin/loglint` |
| `make plugin` | Собрать плагин в `bin/loglint.so` |
| `make test` | Запустить тесты |
| `make lint` | Прогнать golangci-lint |
| `make ci` | Тесты + билд + плагин |
| `make clean` | Удалить `bin/` |