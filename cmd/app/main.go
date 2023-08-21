package main

import (
	"fmt"

	"github.com/reangeline/fiap_teste_ci/internal"
)

func main() {

	teste, err := internal.Somar(10, 10)

	if err != nil {
		fmt.Print(err)
		panic(err)

	}

	print(teste)

}
