package main

import (
	"fmt"
	"sync"
	"time"
)

func imprimir(i int, wg *sync.WaitGroup) {
	d := 1000 * time.Duration(i%5) * time.Millisecond
	time.Sleep(d)
	fmt.Println(fmt.Sprintf("%v numero: %v", d, i))
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go imprimir(i, &wg)
	}
	fmt.Println("Esperanding")
	wg.Wait()
	fmt.Println("Fin!")
}
