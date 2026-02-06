## Описание

Это простой REST API сервер, который демонстрирует:
- Подключение к PostgreSQL через connection pool
- HTTP сервер с базовыми эндпоинтами
- Правильную структуру Go проекта
- Работу с переменными окружения
- Graceful shutdown (корректное завершение работы)
- Структурированное логирование

## Структура проекта

```
project/
├── cmd/
│   └── platform/
│       └── main.go              # Точка входа - здесь запускается приложение
├── internal/
│   ├── app/
│   │   └── app.go               # Основная логика приложения
│   ├── config/
│   │   └── config.go            # Загрузка конфигурации из .env
│   ├── handler/
│   │   └── handler.go           # HTTP handlers (обработчики запросов)
│   └── pkg/
│       ├── httpserver/
│       │   └── server.go        # HTTP сервер
│       └── pgconn/
│           ├── config.go        # Конфигурация PostgreSQL
│           ├── err.go           # Ошибки для работы с БД
│           └── pgconn.go        # Подключение к PostgreSQL
├── .env                         # Конфигурация (НЕ коммитить в git!)
├── .env.example                 # Пример конфигурации
├── .gitignore
├── go.mod                       # Зависимости проекта
├── Makefile                     # Команды для сборки и запуска
└── README.md
```

#### `cmd/platform/main.go`
Точка входа в приложение. Здесь только минимальная логика:
- Инициализация логгера
- Загрузка конфигурации
- Запуск основного приложения

**Почему `cmd`?** - Стандарт Go: здесь хранятся исполняемые файлы (команды).

**Почему `project`?** - Название нашего приложения. Если у вас несколько приложений в одном репозитории, каждое будет в своей папке.

#### `internal/`
**Важно!** Код в `internal/` может использоваться только внутри этого проекта. Это специальная папка Go - другие проекты не смогут импортировать код отсюда.

#### `internal/app/app.go`
Основная бизнес-логика приложения:
- Подключение к базе данных
- Запуск HTTP сервера
- Обработка сигналов завершения (Ctrl+C)
- Graceful shutdown

**Что такое Graceful Shutdown?**
Когда вы нажимаете Ctrl+C, сервер:
1. Перестает принимать новые запросы
2. Завершает обработку текущих запросов
3. Закрывает подключения к БД
4. Только потом останавливается

Это важно для production, чтобы не потерять данные и не оборвать запросы пользователей.

#### `internal/config/config.go`
Загружает настройки из `.env` файла в Go структуры.

**Почему не hardcode значения в коде?**
- Разные настройки для dev/test/production окружений
- Безопасность: пароли не хранятся в коде
- Легко менять без пересборки приложения

#### `internal/handler/handler.go`
HTTP обработчики (handlers) - функции, которые отвечают на запросы:
- `GET /` - информация об API
- `GET /health` - проверка здоровья сервера
- `GET /api/v1/ping` - простой пинг

**Что такое Handler?**
Это функция, которая принимает HTTP запрос и возвращает ответ. Например:
```go
func ping(w http.ResponseWriter, r *http.Request) {
    // r - входящий запрос (request)
    // w - куда писать ответ (response writer)
}
```

**Что такое Middleware?**
Это "обертка" вокруг handler'а. В нашем случае - логирование каждого запроса.

#### `internal/pkg/httpserver/`
Обертка над стандартным `http.Server` с настройками:
- Таймауты (чтобы запросы не висли вечно)
- Graceful shutdown

#### `internal/pkg/pgconn/`
Работа с PostgreSQL:
- **Connection Pool** - набор переиспользуемых подключений к БД
- **Health Check** - проверка, что БД жива

**Что такое Connection Pool?**
Создавать новое подключение к БД для каждого запроса медленно. Pool держит N открытых подключений и переиспользует их. Это намного быстрее!

**Параметры pool:**
- `MaxConns: 20` - максимум 20 одновременных подключений
- Если все заняты, новый запрос подождет освобождения

## Как запустить

### 1. Установите PostgreSQL

Через Docker (самый простой способ):
```bash
docker run --name postgres-dev \
  -e POSTGRES_PASSWORD=1234 \
  -e POSTGRES_DB=platform \
  -p 5432:5432 \
  -d postgres:15-alpine
```

Или установите PostgreSQL на вашу систему.

### 2. Настройте конфигурацию

Скопируйте пример конфигурации:
```bash
cp .env.example .env
```

Отредактируйте `.env` под свои настройки БД.

### 3. Установите зависимости

```bash
go mod download
```

**Что такое `go mod`?**
Система управления зависимостями в Go. Файл `go.mod` содержит список всех библиотек, которые использует проект.

### 4. Запустите приложение

```bash
go run cmd/platform/main.go
```

Или через Makefile:
```bash
make run
```

Сервер запустится на `http://localhost:8080`

## Тестирование API

### Главная страница
```bash
curl http://localhost:8080/
```

Ответ:
```json
{
  "message": "Welcome to Semen Platform API",
  "version": "1.0.0"
}
```

### Health Check
```bash
curl http://localhost:8080/health
```

Ответ:
```json
{
  "status": "ok",
  "database": {
    "status": "healthy",
    "total_conns": 5,
    "idle_conns": 5
  }
}
```

### Ping
```bash
curl http://localhost:8080/api/v1/ping
```

Ответ:
```json
{
  "message": "pong"
}
```

## Основные технологии и концепции

### Go Modules
Система управления зависимостями. Команды:
- `go mod download` - скачать зависимости
- `go mod tidy` - очистить неиспользуемые зависимости
- `go get <package>` - добавить новую зависимость

### pgx/v5
Современный драйвер для PostgreSQL. Быстрее и удобнее стандартного `database/sql`.

**Connection Pool** в pgx:
- Автоматически управляет подключениями
- Переиспользует соединения
- Проверяет здоровье подключений

### Structured Logging (slog)
Вместо простых `fmt.Println()` используем `slog`:
```go
slog.Info("message", "key", "value", "another_key", 123)
```

Выводит JSON:
```json
{"time":"2025-11-05T16:50:56", "level":"INFO", "msg":"message", "key":"value", "another_key":123}
```

**Зачем?** Легко парсить логи в production, фильтровать, анализировать.

### Context
`context.Context` - механизм для передачи отмены операций и таймаутов.

Пример:
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// Операция автоматически отменится через 10 секунд
db.QueryRow(ctx, "SELECT ...")
```

**Зачем?** Если пользователь отменил запрос (закрыл браузер), нет смысла продолжать работу с БД.

### HTTP Server
Стандартная библиотека Go `net/http`:
- `http.NewServeMux()` - роутер (распределяет запросы по handlers)
- `http.HandlerFunc` - тип для handler функций
- `http.ListenAndServe()` - запуск сервера

### Goroutines и Channels
```go
go func() {
    // Эта функция выполняется в отдельном потоке
    serverErrors <- srv.Start()
}()
```

**Goroutine** - легковесный поток. Go может запускать тысячи goroutines.

**Channel** (`chan`) - способ передачи данных между goroutines. Как очередь сообщений.

## Переменные окружения

Настройки в `.env` файле:

| Параметр | Описание | Пример |
|----------|----------|--------|
| `APP_NAME` | Имя приложения | `application` |
| `PUBLIC_API_HOST` | Хост для API | `localhost` |
| `PUBLIC_API_PORT` | Порт для API | `8080` |
| `PG_HOST` | Хост PostgreSQL | `localhost` |
| `PG_PORT` | Порт PostgreSQL | `5432` |
| `PG_USER` | Пользователь БД | `postgres` |
| `PG_PASS` | Пароль БД | `1234` |
| `PG_DB` | Имя базы данных | `platform` |
| `PG_MAX_CONN` | Макс. подключений | `20` |

## Полезные команды

```bash
# Запуск
go run cmd/platform/main.go

# Сборка бинарника
go build -o bin/platform cmd/platform/main.go

# Запуск бинарника
./bin/platform

# Форматирование кода
go fmt ./...

# Проверка кода
go vet ./...

# Установка зависимостей
go mod download

# Очистка зависимостей
go mod tidy
```
# main_project
