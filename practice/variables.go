package main

import "fmt"

// Глобальные: только декларации var/const
var x int = 42
var y int // zero value: 0

const Pi = 3.14159
const (
	Sunday  = iota // 0
	Monday         // 1
	Tuesday        // 2
)

func sumEven(max int) int {
	sum := 0
	for i := 0; i <= max; i += 2 {
		sum += i
	}
	return sum
}
func main1() {
	// Здесь statements и :=
	y = 10 // присвоение ок

	z := 3.14       // float64
	name := "Гофер" // string

	a, b := 1, "два"
	fmt.Println("sumEven", sumEven(100))
	fmt.Println(x, y, z, name, a, b)
	fmt.Printf("Pi = %.5f, Sunday=%d\n", Pi, Sunday)
}
