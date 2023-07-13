package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// a função gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "busca ips e nome de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca ips de endereços na internet",
			Flags: flags,

			Action: buscarIps,
		},

		{
			Name:   "servidores",
			Usage:  "busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	//funçção do pacote Net
	ips, erro := net.LookupHost(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}
