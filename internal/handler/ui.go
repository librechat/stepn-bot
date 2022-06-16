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
			"https://www.food365.shop/wp-content/uploads/2020/04/%E9%9F%AD%E8%8F%9C-500x500-1.jpg",
			"Welcome to Stepn together", "A small walk worth a coffee break",
			&linebot.PostbackAction{
				Label:       "GST-SOL 懶人匯率",
				DisplayText: "GST-SOL 懶人匯率",
				Data:        "lazy",
			},
			&linebot.MessageAction{
				Label: "HELP",
				Text:  "help",
			},
			&linebot.MessageAction{
				Label: "Free Msg Quota",
				Text:  "quota",
			},
		),
	)
}
