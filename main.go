package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("ponto de partida")

	aplicação := app.Gerar()
	if erro := aplicação.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
	time.Sleep(50 * time.Second)
}
