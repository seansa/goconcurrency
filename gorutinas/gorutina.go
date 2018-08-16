package main

import (
	"fmt"
	"time"
)

const ESPERA = 9 * time.Millisecond

func imprimir(i int) {
	d := 10 * time.Duration(i%2) * time.Millisecond
	time.Sleep(d)
	fmt.Println(fmt.Sprintf("%v numero: %v", d, i))
}

func main() {
	for i := 1; i < 11; i++ {
		go imprimir(i)
	}
	time.Sleep(ESPERA)
}
