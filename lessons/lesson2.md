# Урок 2: Переменные, типы данных и базовые конструкции

## Теория

Go — строго типизированный, компилируемый язык. Нет `undefined/null` как в JS (есть `nil` только для интерфейсов/указателей/каналов/slice/map). Ошибки типов ловятся на компиляции.

### 1. Переменные (аналогия с JS `let/const`)
- `var имя тип = значение` (явная).
- `var имя тип` (zero value).
- `:=` (short declaration, только в func).
- `const` (compile-time).

**Zero values**: int=0, string="", bool=false, slice/nil.

**iota** для enum-like.

```/dev/null/variables.go#L1-35
package main

import "fmt"

var x int = 42
var y int  // 0

const Pi = 3.14159
const (
    Sunday = iota
    Monday
)

func main() {
    y = 10
    z := 3.14
    fmt.Println(x, y, z)
}
```

### 2. Типы данных
- Числа: `int`, `int64`, `float64`, `byte` (uint8), `rune` (Unicode).
- Строки: immutable `string` (UTF-8 bytes).

```/dev/null/types.go#L1-15
s := "hello"
b := []byte(s)
r := []rune(s)
```

### 3. Управление потоком
- `if cond {}` (no parens, init: `if n:=9; n<10 {}`).
- `for` only: `for i:=0; i<10; i++ {}`, `while`: `for cond {}`, `foreach`: `for _, v := range slice {}`.
- `switch` no fallthrough.

```/dev/null/control.go#L1-20
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

### 4. Enum паттерны (нет строковых enum)
- Slice: `dayNames[day]`.
- Map.
- Type + `String() string`.

```/dev/null/enum.go#L1-30
type Weekday int
const ( Sunday Weekday = iota ... )
func (d Weekday) String() string { ... }
```

## Задания (выполнены)
1. Создай `variables.go`, запусти, покажи вывод.
2. Добавь ошибку типов (string to int) — compile error.
3. `func sumEven(max int) int` с for (чётные 1..max), вызов `sumEven(100)` = 2550.

**Статус**: ✅ Выполнено (подтверждено пользователем).