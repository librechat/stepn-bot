package handler

import (
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Handler struct {
	Bot       *linebot.Client
	ActionMap map[string]func(...string) linebot.SendingMessage
}

var actionMap = map[string]func(...string) linebot.SendingMessage{
	"help":     text(help),
	"lazy":     text(lazy),
	"price":    text(price),
	"prices":   text(richPrice),
	"exchange": text(exchange),
	"menu":     menu,
}

func New(bot *linebot.Client) *Handler {
	h := &Handler{
		Bot: bot,
	}
	actionMap["quota"] = text(h.quota)
	return h
}

func (h *Handler) Handle(event *linebot.Event) {
	reply := handle(event)
	if reply != nil {
		if _, err := h.Bot.ReplyMessage(event.ReplyToken, reply).Do(); err != nil {
			log.Print(err)
		}
	}
}

func handle(event *linebot.Event) linebot.SendingMessage {
	switch event.Type {
	case linebot.EventTypeMessage:
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			cmd, params := parseCmd(message.Text)
			if h, ok := actionMap[cmd]; ok {
				return h(params...)
			}
		}
	case linebot.EventTypePostback:
		cmd, params := parseCmd(event.Postback.Data)
		if h, ok := actionMap[cmd]; ok {
			return h(params...)
		}
	}
	return actionMap["menu"]()
}

func parseCmd(line string) (string, []string) {
	params := strings.Split(line, " ")
	return params[0], params[1:]
}
