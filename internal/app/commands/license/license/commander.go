package license

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/istrel/bot/internal/app/path"
	"github.com/istrel/bot/internal/service/license/license"
)

type LicenseLicenseCommander struct {
	bot            *tgbotapi.BotAPI
	licenseService *license.LicenseService
}

func NewLicenseLicenseCommander(
	bot *tgbotapi.BotAPI,
) *LicenseLicenseCommander {
	licenseService := license.NewDummyLicenseService()

	return &LicenseLicenseCommander{
		bot:            bot,
		licenseService: licenseService,
	}
}

func (c *LicenseLicenseCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LicenseLicenseCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LicenseLicenseCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
