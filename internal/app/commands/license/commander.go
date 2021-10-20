package license

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/istrel/bot/internal/app/commands/license/license"
	"github.com/istrel/bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type LicenseCommander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander Commander
}

func NewLicenseCommander(
	bot *tgbotapi.BotAPI,
) *LicenseCommander {
	return &LicenseCommander{
		bot: bot,
		// subdomainCommander
		subdomainCommander: license.NewLicenseLicenseCommander(bot),
	}
}

func (c *LicenseCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "license":
		c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LicenseCommander.HandleCallback: unknown license - %s", callbackPath.Subdomain)
	}
}

func (c *LicenseCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "license":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("LicenseCommander.HandleCommand: unknown license - %s", commandPath.Subdomain)
	}
}
