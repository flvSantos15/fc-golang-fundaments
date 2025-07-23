package generics

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return  soma
}

func Soma[T comparable] (number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	
	return number2
}

func Exec() {
	var x, y, z MyNumber
	x = 1
	y = 2
	z = 3

	fmt.Println(SomaGenerica(map[string]MyNumber{"a":x, "b":y, "c": z}))
	
	fmt.Println(SomaGenerica(map[string]int64{"a":1, "b":2, "c": 3}))
	fmt.Println(SomaGenerica(map[string]float64{"a":1.1, "b":22.2, "c": 3.2}))
}