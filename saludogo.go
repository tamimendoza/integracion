package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tamimendoza/saludos"
)

func main() {
	canal := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			canal <- i
			time.Sleep(time.Second)
		}
		close(canal)
	}()

	go func() {
		for numero := range canal {
			fmt.Println("Recibí el número: ", numero)
		}
	}()

	time.Sleep(5 * time.Second)
}

func imprimir(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, ":", s)
	}
}

func main_gorutinas() {
	go imprimir("mundo")
	imprimir("hola")
}

func main_integracion() {
	log.SetPrefix("saludos: ")
	log.SetFlags(0)

	nombres := []string{"Gladys", "Gustavo", "Abby"}

	mensajes, err := saludos.SaludosATodos(nombres)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mensajes)
}
