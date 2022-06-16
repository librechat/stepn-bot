package httpserver

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/librechat/stepn-bot/internal/handler"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type server struct {
	Bot    *linebot.Client
	Router *gin.Engine
}

func New(bot *linebot.Client) *server {
	return &server{
		Bot: bot,
	}
}

func (s *server) Route() *gin.Engine {
	if s.Router != nil {
		return s.Router
	}

	s.Router = gin.Default()
	s.Router.Use(s.ValidateRequest)
	//s.Router.GET("/callback", s.Echo)
	s.Router.POST("/callback", s.Handle)

	return s.Router
}

func (s *server) Handle(c *gin.Context) {
	var events []*linebot.Event
	if val, ok := c.Get("events"); ok {
		events, _ = val.([]*linebot.Event)
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				msg := handler.Handle(message.Text)

				// GetMessageQuota: Get how many remain free tier push message quota you still have this month. (maximum 500)
				quota, err := s.Bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				// message.ID: Msg unique ID
				// message.Text: Msg text
				msg += fmt.Sprintf("remain Msg Count = %d", quota.Value)
				if _, err = s.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
					log.Print(err)
				}
			}
		case linebot.EventTypePostback:
		}
	}
}
