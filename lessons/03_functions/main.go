package main

import "fmt"

func avg(nums ...float32) {
	for i, n := range nums {
		fmt.Println(i, n)
	}
}
func main() {
	avg(1, 2, 3, 4, 5)
}
