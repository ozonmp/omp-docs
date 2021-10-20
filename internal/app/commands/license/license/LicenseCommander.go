package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/istrel/bot/internal/service/license/license"
)

type LicenseCommander struct {
	bot            *tgbotapi.BotAPI
	licenseService *license.LicenseService
}

func NewLicenseCommander(
	bot *tgbotapi.BotAPI,
	licenseService *license.LicenseService,
) *LicenseCommander {
	return &LicenseCommander{
		bot:            bot,
		licenseService: licenseService,
	}
}
