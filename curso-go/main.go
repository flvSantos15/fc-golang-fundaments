package main

import (
	"fmt"

	math "github.com/flvSantos15/fc-golang-fundaments/math"
)

func main() {
	soma()
}

func soma() {
	resultado := math.Soma(1, 2)
	fmt.Println(resultado)
}

