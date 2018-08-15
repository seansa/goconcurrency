package main

import (
	"fmt"
	"time"
)

func imprimir(i int, sinc chan string) {
	d := 10 * time.Duration(i%2) * time.Millisecond
	time.Sleep(d)
	sinc <- fmt.Sprintf("%v numero: %v", d, i)
}

func main() {
	var sincronizador1 = make(chan string)
	var sincronizador2 = make(chan string)
	var sincronizador3 = make(chan string)
	var sincronizador4 = make(chan string)
	var sincronizador5 = make(chan string)

	go imprimir(1, sincronizador1)
	go imprimir(2, sincronizador2)
	go imprimir(3, sincronizador3)
	go imprimir(4, sincronizador4)
	go imprimir(5, sincronizador5)

	for i := 0; i < 5; i++ {
		select {
		case i1 := <-sincronizador1:
			fmt.Println(i1)
		case i2 := <-sincronizador2:
			fmt.Println(i2)
		case i3 := <-sincronizador3:
			fmt.Println(i3)
		case i4 := <-sincronizador4:
			fmt.Println(i4)
		case i5 := <-sincronizador5:
			fmt.Println(i5)
		}
	}
	fmt.Println("Fin")
}
