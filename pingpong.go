package main

import (
	"fmt"
	"time"
)

func pingar(c chan string, c2 chan string) {
	for i := 0; ; i++ {
		c <- "ping" // usado para enviar e receber mensagem pelo canal
		c2 <- "pong"
	}
}

func imprimir(c chan string, c2 chan string) {
	for {
		msg := <-c
		msg2 := <-c2
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
		fmt.Println(msg2)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	var c2 chan string = make(chan string)

	go pingar(c, c2)
	go imprimir(c, c2)

	var entrada string
	fmt.Scanln(&entrada)
}
