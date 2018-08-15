package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Mugir struct {
	Vaca      int
	Sonido    string
	Ordeñando chan bool
}

func main() {
	runGranja(2)
	fmt.Println("Listo!!")
}

// La granja comienza con mugidos de vacas que quieren ser ordeñadas
func runGranja(vacas int) {
	canalGranjero := make(chan string)
	canalVacas := make(chan Mugir)
	for numVaca := 0; numVaca < vacas; numVaca++ {
		go vaca(numVaca+1, canalVacas)
	}
	go granjero(vacas, canalVacas, canalGranjero)
	elGranjeroDice := <-canalGranjero
	if elGranjeroDice == "yey!" {
		fmt.Println("Todas las vacas son felices.")
	}
}

// Una vaca mugirá hasta que sea ordeñada
func vaca(num int, canalVacas chan Mugir) {
	canal := make(chan bool)
	for {
		select {
		case canalVacas <- Mugir{num, "muu", canal}:
			fmt.Println("Vaca numero", num, "mugiendo a travez del canalVacas")
			<-canal
			fmt.Println("Vaca numero", num, "está siendo ordeñada y deja de mugir")
			canalVacas <- Mugir{num, "moohh", nil}
			fmt.Println("Vaca numero", num, "mugio una última vez como alivio")
			return
		default:
			fmt.Println("Vaca numero", num, "mugiendo a travez del canalVacas y fue ignorada")
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		}
	}
}

// El granjero quiere encender todos los tubos de leche para detener el mugido
func granjero(vacas int, canalVacas chan Mugir, canalGranjero chan string) {
	fmt.Println("El granjero comienza a escuchar el canalVacas")
	for vacasSinAlivio := vacas; vacasSinAlivio > 0; {
		muu := <-canalVacas
		if muu.Sonido == "moohh" {
			fmt.Println("El granjero escuchó un mugido de alivio de vaca numero", muu.Vaca)
			vacasSinAlivio--
		} else {
			fmt.Println("El granjero oyó un", muu.Sonido, "de la vaca numero", muu.Vaca)
			time.Sleep(2e9)
			fmt.Println("El granjero enciende la máquina de ordeñar con la vaca numero", muu.Vaca)
			muu.Ordeñando <- true
		}
	}
	fmt.Println("El granjero ya no oye ni un solo muu. ¡Todo listo!")
	canalGranjero <- "yey!"
}
