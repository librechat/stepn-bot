package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/librechat/stepn-bot/internal/httpserver"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	bot, err := linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))
	log.Println("Bot:", bot, " err:", err)

	http.Handle("/", httpserver.New(bot).Route())

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}
