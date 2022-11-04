package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()

}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for acess to telegeram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal()
	}

	return *token
}
