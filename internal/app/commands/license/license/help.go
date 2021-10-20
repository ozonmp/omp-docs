package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LicenseCommander) HelpLicenseLicense(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__license__license - help\n"+
			"/get__license__license - get a entity\n"+
			"/list__license__license - list licenses\n"+
			"/delete__license__license - delete an existing entity\n"+
			"\n"+
			"/new__license__license - create a new entity\n"+
			"/edit__license__license - edit an entity\n",
	)

	c.bot.Send(msg)
}
