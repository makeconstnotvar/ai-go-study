# Урок 1: Основы Go — Установка, Hello World, базовые команды

## Введение
Go (Golang) — компилируемый, статически типизированный язык с garbage collection. Аналогия с JS: как Node.js, но типы явные (нет dynamic typing), компиляция в бинарник (быстрее runtime), встроенная concurrency (goroutines > async/await).

## 1. Установка
- Скачай с [go.dev/dl](https://go.dev/dl/) (stable для macOS/Linux/Windows).
- Проверь: `go version` (должен показать 1.21+).
- Отличие от Node: один бинарник, нет npm/global deps (пакеты в `go mod`).

## 2. Структура проекта
- Директория проекта.
- `go.mod` — манифест (как `package.json`).
- `main.go` — entry point (package main).

## Пример: Hello World
Создай папку `hello-go`, файл `main.go`:

```/dev/null/hello-go/main.go#L1-10
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

**Команды** (в папке hello-go):
- `go mod init hello-go` — создаёт `go.mod`.
- `go run main.go` — компиляция+запуск (как `node main.js`).
- `go build` — бинарник `hello-go` (без .exe на Unix, запуск `./hello-go`). Нормально!

**Вывод**: `Hello, Go!`

## Отличия от JS
| Go | JS |
|----|----|
| Статические типы (compile errors) | Dynamic (runtime) |
| `go build` → standalone binary | `node app.js` → runtime deps |
| Нет `undefined/null` (zero values) | undefined/null/NaN |

## Задания (выполнено)
1. Установи Go, проверь `go version`.
2. Создай `hello-go/main.go`, `go mod init`, `go run` — покажи вывод.
3. `go build`, запусти бинарник `./hello-go`.

**Статус**: ✅ Выполнено (подтверждено: "Hello, Go!", бинарник hello-go).

## Следующий: Урок 2 — Переменные и типы.