package main1

import (
	"errors"
	"fmt"
)

// Базовая
func add(a, b int) int {
	return a + b
}

// Множественные возвраты + error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

// Variadic (...int → []int внутри)
func sumEven(nums ...int) int {
	sum := 0
	for _, n := range nums {
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}

// Closure factory
func power(base int) func(int) int {
	return func(exp int) int {
		result := 1
		for i := 0; i < exp; i++ {
			result *= base
		}
		return result
	}
}

func main() {
	defer fmt.Println("Конец программы") // defer всегда

	fmt.Println("add(3,4):", add(3, 4)) // 7

	res, err := divide(10, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("divide:", res) // ~3.333
	}

	fmt.Println("sumEven(1,2,3,4,5):", sumEven(1, 2, 3, 4, 5))   // 6 (2+4)
	fmt.Println("sumEven из slice:", sumEven([]int{6, 7, 8}...)) // 14

	pow2 := power(2)
	fmt.Println("2^3 =", pow2(3)) // 8
}
