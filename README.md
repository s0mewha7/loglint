# loglint (Selectel Internship Test Task)

Линтер для Go, анализирует лог-записи в коде и проверяет их соответствие правилам стиля и безопасности.

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
│   └── main.go              # точка входа, запуск анализатора
├── internal/
│   ├── analyzer/
│   │   └── analyzer.go      # обход AST дерева, вызов правил
│   ├── rules/
│   │   ├── rules.go         # реализация правил проверки
│   │   └── rules_test.go    # unit-тесты для каждого правила
│   └── utils/
│       └── utils.go         # определение лог-вызова, извлечение сообщения
├── testdata/src/
│   ├── slog/
│   │   └── slog.go          # тестовые случаи для log/slog
│   └── uberzap/
│       └── zap.go           # тестовые случаи для go.uber.org/zap
├── .github/
│   └── workflows/
│       └── ci.yml           # CI/CD pipeline
├── .gitignore
├── .golangci.yml
├── Makefile
└── README.md
```

## Требования

- Go 1.22+

## Установка
```bash
git clone https://github.com/s0mewha7/loglint
cd loglint
go mod tidy
```

## Сборка и запуск

Собрать бинарник и прогнать линтер на тестовых файлах:
```bash
make all
```

Или по шагам:
```bash
make build   # собрать бинарник ./loglint
make run     # запустить линтер на testdata
make test    # запустить unit-тесты
make clean   # удалить бинарник
```

Без make:
```bash
go build -o loglint ./cmd/loglint
go vet -vettool=./loglint ./testdata/src/slog/slog.go
go vet -vettool=./loglint ./testdata/src/uberzap/zap.go
```

## Тесты

Unit-тесты покрывают каждое правило отдельно:

- `TestLowercase` — проверка заглавной буквы в начале сообщения
- `TestEnglish` — проверка на английский язык
- `TestSpecialChars` — проверка на спецсимволы и эмодзи
- `TestSecrets` — проверка на чувствительные данные

Запуск:
```bash
make test
```

Вывод:
```
ok  github.com/s0mewha7/loglint/internal/rules  0.002s
```

## CI/CD

Проект настроен на автоматическую сборку и тестирование через GitHub Actions.

Пайплайн запускается при каждом пуше и pull request в ветку `main` и выполняет три шага:

1. `make build` — сборка бинарника
2. `make test` — запуск unit-тестов
3. `make run` — прогон линтера на testdata

Конфиг: `.github/workflows/ci.yml`

Статус пайплайна виден во вкладке **Actions** на странице репозитория.

## Примеры
```go
// ❌ не пройдёт

// log/slog
slog.Info("Starting server on port 8080")  // заглавная буква
slog.Info("запуск сервера")                // не английский
slog.Warn("server started!🚀")             // спецсимволы
slog.Info("user login: " + password)       // чувствительные данные

// go.uber.org/zap
logger.Info("Starting server")             // заглавная буква
logger.Error("ошибка подключения")         // не английский
logger.Warn("server started!🚀")           // спецсимволы
logger.Debug("auth: " + token)             // чувствительные данные

// ✅ пройдёт

slog.Info("starting server on port 8080")
slog.Error("failed to connect to database")
slog.Info("user authenticated successfully")

logger.Info("starting server on port 8080")
logger.Error("failed to connect to database")
logger.Info("user authenticated successfully")
```

Вывод линтера:
```
./testdata/src/slog/slog.go:8:2: log message must start with lowercase
./testdata/src/slog/slog.go:12:2: log message must be in english
./testdata/src/slog/slog.go:16:2: log message must not contain special characters
./testdata/src/slog/slog.go:21:2: possible sensitive data in logs: password

./testdata/src/uberzap/zap.go:9:2: log message must start with lowercase
./testdata/src/uberzap/zap.go:13:2: log message must be in english
./testdata/src/uberzap/zap.go:17:2: log message must not contain special characters
./testdata/src/uberzap/zap.go:22:2: possible sensitive data in logs: token
```
