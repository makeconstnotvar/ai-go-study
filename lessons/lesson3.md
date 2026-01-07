# Урок 3: Функции в Go

## Введение
Функции — основной строительный блок. Аналогия с JS: функции first-class, но в Go **явная сигнатура** (типы везде), **множественные возвраты** (как [value, error]), нет default params, closures как factory funcs, `defer` как finally (LIFO стек).

Отличия от JS:
- Нет arrow funcs (все func).
- Нет hoisting.
- `defer` отложенный вызов (LIFO: последний defer первый).

## 1. Базовые функции
```go
func add(a int, b int) int {
    return a + b
}
```
Вызов: `add(1,2)`.

## 2. Variadic (`...`)
Как JS rest: `func sum(nums ...int) int { for _, n := range nums { sum += n } }`
Вызов: `sum(1,2,3)` или `sum(1,2,3...)`.

## 3. Множественные возвраты
Стандарт: `func divide(a, b float64) (float64, error) { if b==0 { return 0, errors.New(\"div0\") } return a/b, nil }`
Разыменование: `q, err := divide(10,3); if err != nil { ... }`.

## 4. Closures (анонимные функции)
Factory: `func multiplier(factor int) func(int) int { return func(x int) int { return x * factor } }`
Исп: `double := multiplier(2); double(5)`.

## 5. Defer (LIFO)
`defer fmt.Println(\"world\"); defer fmt.Println(\"hello\")` → hello\nworld.
Полезно для cleanup: file.Close() в defer.

## Примеры в коде
См. hello-go/functions.go (добавь/обнови).

## Задания (ожидаю код + вывод)
1. `func avg(nums ...float64) float64` — среднее.
2. `func safeDivide(a, b int) (int, error)` — деление с err.
3. Generator: `func gen(n int) func() int { i:=0; return func() int { defer func(){i++;}(); return i } }` — тест.
4. Функция с 3 defer: покажи порядок.

Выполни в functions.go, `go run functions.go`, покажи вывод + код.

**Статус**: Ожидаю выполнения.

## Следующий: Урок 4 — Структуры и методы.
