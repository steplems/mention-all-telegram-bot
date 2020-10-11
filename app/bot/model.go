package bot

import (
	"github.com/akali/steplems-bot/app/commands"
	"github.com/akali/steplems-bot/app/database"
	"github.com/akali/steplems-bot/app/logger"
	"github.com/go-bongo/bongo"
	tbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

type (
	// Bot is a wrapper for tbot.BotAPI that stricts and simplifies
	// its functionality.
	Bot struct {
		RunBotRepo
		RecordMessageRepo
		api      *tbot.BotAPI
		commands commands.CallbackMap
		Database *database.Database
	}
)

var (
	log = logger.Factory.NewLogger("bot")
)

// NewBot initializes bot api and returns a new *Bot.
func NewBot(token string, commands commands.CallbackMap, config *bongo.Config) (*Bot, error) {
	api, err := tbot.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	databaseConfig := &database.Database{
		Config: config,
	}

	return &Bot{api: api, commands: commands, Database: databaseConfig}, nil
}
