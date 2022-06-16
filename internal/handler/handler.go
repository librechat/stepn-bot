package handler

import (
	"fmt"

	"github.com/librechat/stepn-bot/internal/crypto"
)

func Handle(cmd string) string {
	switch cmd {
	case "help":
		return help()
	case "price":
		return price()
	}
	return ""
}

func help() string {
	return "1. help\n2. price - to show GST and SOL exchanges"
}

func price() string {
	// show price
	sol := crypto.GetCoinData(crypto.SOL).MarketData.CurrentPrice["usd"]     // 1 sol = ? usd
	gst := crypto.GetCoinData(crypto.GST_SOL).MarketData.CurrentPrice["usd"] // 1 gst = ? usd

	ex := gst / sol // 1 gst = ? sol

	return fmt.Sprintf("1 SOL = %.5f USD\n1 GST = %.5f USD\n1 GST = %.5f SOL\n", sol, gst, ex)
}
