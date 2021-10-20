package license

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LicenseLicenseCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	license, err := c.licenseService.Get(idx)
	if err != nil {
		log.Printf("fail to get license with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		license.Title,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LicenseLicenseCommander.Get: error sending reply message to chat - %v", err)
	}
}
