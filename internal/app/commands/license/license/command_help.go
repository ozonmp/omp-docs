package license

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LicenseLicenseCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list licenses",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LicenseLicenseCommander.Help: error sending reply message to chat - %v", err)
	}
}
