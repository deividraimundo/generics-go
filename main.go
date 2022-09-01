package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// func SomaGenerica[T int64 | float64](m map[string]T) T {
// 	var soma T
// 	for _, v := range m {
// 		soma += v
// 	}

// 	return soma
// }

// Criando um tipo
type Number interface {
	int64 | float64
}

// *** Tipos com Generics ***
func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

// *** Comparação com Generics ***
func Soma[T comparable](number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}

	return number2
}

// *** Comparação com biblioteca constraint de Generics ***
func M[T constraints.Ordered](number1 T, number2 T) T {
	if number1 > number2 {
		return number1
	}

	return number2
}

// *** Concatenação com Generics ***
type MyString string

func (m MyString) String() string {
	return string(m)
}

type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}

func main() {

	fmt.Println(concat([]MyString{"a", "b", "c"}))

	fmt.Println(SomaGenerica(map[string]int64{"a": 1, "b": 2, "c": 3}))
	fmt.Println(SomaGenerica(map[string]float64{"a": 1.1, "b": 2.5, "c": 3.6}))
}
