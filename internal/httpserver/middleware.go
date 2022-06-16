package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (s *server) ValidateRequest(c *gin.Context) {
	events, err := s.Handler.Bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Set("events", events)

	c.Next()
}
