package httpserver

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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
	s.Router.POST("/callback", s.EchoLine)

	return s.Router
}

func (s *server) EchoLine(c *gin.Context) {
	var events []*linebot.Event
	if val, ok := c.Get("events"); ok {
		events, _ = val.([]*linebot.Event)
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// Handle only on text message
			case *linebot.TextMessage:
				// GetMessageQuota: Get how many remain free tier push message quota you still have this month. (maximum 500)
				quota, err := s.Bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}
				// message.ID: Msg unique ID
				// message.Text: Msg text
				if _, err = s.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("msg ID:"+message.ID+":"+"Get:"+message.Text+" , \n OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
