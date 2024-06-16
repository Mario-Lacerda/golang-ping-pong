package main

import (
	"fmt"
	"time"
)

func ping(canal chan string) {
	for i := 0; ; i++ {
		canal <- "ping" // envia a mensagem para o "canal"
	}
}

func pong(canal chan string) {
	for i := 0; ; i++ {
		canal <- "pong" // envia a mensagem para o "canal"
	}
}

func imprimir(canal chan string) {
	for {
		mensagem := <-canal
		fmt.Println(mensagem)
		time.Sleep(time.Second * 1) 		
	}
}

func main() {

	var canal chan string = make(chan string)
	go ping(canal)
	go imprimir(canal)
	go pong(canal)

	var aguardaUsuario string 
	fmt.Scanln(&aguardaUsuario)

}
