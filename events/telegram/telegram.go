package telegram

import "github.com/Umk1nus/bot-go-for-tg/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}
