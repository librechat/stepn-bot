package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/librechat/stepn-bot/internal/handler"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type server struct {
	Handler *handler.Handler
	Router  *gin.Engine
}

func New(bot *linebot.Client) *server {
	return &server{
		Handler: handler.New(bot),
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
		s.Handler.Handle(event)
	}
}
