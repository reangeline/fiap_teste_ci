package internal

import "fmt"

func Somar(a, b float32) (float32, error) {
	fmt.Println(a + b)
	return a + b, nil

}
