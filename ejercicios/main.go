package main

import (
	"errors"
	"fmt"
)

func division(num1 float32, num2 float32) (float32, error) {
	var resultado float32
	var err error = nil // error nulo

	if num2 == 0 {
		err = errors.New("hubo un error, estas diviidendo por 0")
		return 0, err

	}

	resultado = num1 / num2

	return resultado, err
}

func main() {
	var div float32
	var err error
	div, err = division(10, 0)

	if err != nil {
		fmt.Println(err.Error())

	}

	fmt.Println(div)
}
