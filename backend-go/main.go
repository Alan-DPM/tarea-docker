package main

import "fmt"

func Sumar(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("La suma de 2 + 3 en Go es:", Sumar(2, 3))
}
