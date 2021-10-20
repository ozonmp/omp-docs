package commands

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LicenseCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote 123: "+inputMessage.Text)

	c.bot.Send(msg)
}

func (c *LicenseCommander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleMessage(update.Message)
	}
}

func (c *LicenseCommander) handleCallback(callback *tgbotapi.CallbackQuery) {
	parsedData := CommandData{}
	json.Unmarshal([]byte(callback.Data), &parsedData)
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed license: %+v\n", parsedData),
	)
	c.bot.Send(msg)
}

func (c *LicenseCommander) handleMessage(msg *tgbotapi.Message) {
	switch msg.Command() {
	case "help__license__license":
		c.HelpLicenseLicense(msg)
	case "list__license__license":
		c.List(msg)
	case "get__license__license":
		c.Get(msg)
	case "delete__license__license":
		c.Get(msg)
	case "new__license__license":
		c.Get(msg)
	case "edit__license__license":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
