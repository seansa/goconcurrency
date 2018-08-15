package ejemplos

import (
	"fmt"
)

func Ejemplo1() {
	ch := make(chan string, 3)
	ch <- "valor 1"
	ch <- "valor 2"
	ch <- "valor 3"
	fmt.Println("capacidad: ", cap(ch))
	fmt.Println("length: ", len(ch))
	fmt.Println("ch: ", <-ch)
	fmt.Println("length: ", len(ch))
}
