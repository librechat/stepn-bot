package handler

import (
	"fmt"
	"log"
)

func (h *Handler) quota(s ...string) string {
	// GetMessageQuota: Get how many remain free tier push message quota you still have this month. (maximum 500)
	quota, err := h.Bot.GetMessageQuota().Do()
	if err != nil {
		log.Println("quota err:", err)
	}
	return fmt.Sprintf("Remain free msg count = %d", quota.Value)
}
