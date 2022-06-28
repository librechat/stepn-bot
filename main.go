package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/librechat/stepn-bot/internal/httpserver"
	"github.com/librechat/stepn-bot/richmenu"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	bot, err := linebot.New(secret())
	log.Println("Bot:", bot, " err:", err)

	if err := richmenu.CreateRichMenu(bot); err != nil {
		println(err.Error())
	}

	http.Handle("/", httpserver.New(bot).Route())

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func secret() (string, string) {
	return os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN")
}
