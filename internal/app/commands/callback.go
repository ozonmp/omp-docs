package commands

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Callback(callback *tgbotapi.CallbackQuery) {
	parsedData := CommandData{}
	json.Unmarshal([]byte(callback.Data), &parsedData)
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %+v\n", parsedData),
	)
	c.bot.Send(msg)
}
