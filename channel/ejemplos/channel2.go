package ejemplos

import (
	"fmt"
	"time"
)

func imprimir(i int, sinc chan string) {
	d := 10 * time.Duration(i%2) * time.Millisecond
	time.Sleep(d)
	sinc <- fmt.Sprintf("%v numero: %v", d, i)
}

func Ejemplo2() {
	var sincronizador = make(chan string)
	for i := 1; i <= 10; i++ {
		go imprimir(i, sincronizador)
	}
	for i := 1; i <= 10; i++ {
		resp := <-sincronizador
		fmt.Println(resp)
	}
	fmt.Println("Fin!")
}
