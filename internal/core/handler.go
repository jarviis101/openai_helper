package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"helper_openai_bot/internal/core/command"
	"helper_openai_bot/internal/core/text"
)

type Handler interface {
	Handle(update tgbotapi.Update)
}

type handler struct {
	commandHandler command.Handler
	textHandler    text.Handler
}

func CreateHandler(commandHandler command.Handler, textHandler text.Handler) Handler {
	return &handler{
		commandHandler,
		textHandler,
	}
}

func (h *handler) Handle(update tgbotapi.Update) {
	if update.Message.IsCommand() {
		h.commandHandler.Handle(update)
		return
	}

	h.textHandler.Handle(update)
}
