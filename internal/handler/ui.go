package handler

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func text(reply func(...string) string) func(...string) linebot.SendingMessage {
	return func(p ...string) linebot.SendingMessage {
		return linebot.NewTextMessage(reply(p...))
	}
}

func menu(p ...string) linebot.SendingMessage {
	return linebot.NewTemplateMessage(
		"menu unavailable",
		linebot.NewButtonsTemplate(
			"data:image/jpeg;base64,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",
			"Welcome to Stepn together", "A small walk worth a coffee break",
			&linebot.PostbackAction{
				Label: string(linebot.ActionTypePostback),
				Text:  "GST-SOL 懶人匯率",
				Data:  "lazy",
			},
			&linebot.PostbackAction{
				Label: string(linebot.ActionTypePostback),
				Text:  "help",
				Data:  "help",
			},
			&linebot.PostbackAction{
				Label: string(linebot.ActionTypePostback),
				Text:  "quota",
				Data:  "quota",
			},
		),
	)
}
