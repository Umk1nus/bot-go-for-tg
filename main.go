package main

import (
	"flag"
	"github.com/Umk1nus/bot-go-for-tg/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal()
	}

	return *token
}
