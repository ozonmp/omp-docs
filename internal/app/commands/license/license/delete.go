package commands

import (
	"encoding/json"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LicenseCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	license, err := c.licenseService.Delete(idx)
	if err != nil {
		log.Printf("fail to delete license with idx %d: %v", idx, err)
		return
	}

	msgText := "can not delete "
	if license {
		msgText = "successfully delete "
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgText+string(rune(idx)),
	)

	c.bot.Send(msg)

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

}
