package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Print("Hello ")
	fmt.Print("world!\n")
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}

	n := "nesto"
	fmt.Println(n[0])
	fmt.Println(len(n))

	for i, c := range n {
		fmt.Println(i, string(c))
	}
	for i := 0; i < len(n); i++ {
		fmt.Println(string(n[i]))
	}

	fmt.Println("Upisite prvi broj")
	var x int
	fmt.Scanf("%d", &x)
	fmt.Println("Upisite drugi broj")
	var y int
	fmt.Scanf("%d", &y)

	fmt.Println("Funkcija zbrajanja za ", x, " i ", y, ": ", add(x, y))

	fmt.Println("Testing output!")

}

func add(a int, b int) int {
	sum := a + b
	return sum
}
