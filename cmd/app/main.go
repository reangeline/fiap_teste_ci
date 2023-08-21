package main

import "github.com/reangeline/fiap_teste_ci/internal"

func main() {

	teste, err := internal.Somar(10, 10)

	if err != nil {
		panic(err)

	}

	print(teste)

}
