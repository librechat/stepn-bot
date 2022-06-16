package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/librechat/stepn-bot/internal/crypto"
)

func HandleCmd(line string) string {
	params := strings.Split(line, " ")
	switch params[0] {
	case "help":
		return help()
	case "lazy":
		return lazy()
	case "price":
		if len(params) < 2 {
			return "please input coin name"
		}
		return price(params[1])
	case "exchange":
		if len(params) < 4 {
			return "please input valid params - exchange (currency val) (from which coin) (to which coin)"
		}
		cnt, _ := strconv.ParseFloat(params[1], 64)
		return exchange(cnt, params[2], params[3])
	}
	return ""
}

func help() string {
	return "Hi you may use below commands :-)\n1. help\n2. lazy - to show GST and SOL prices\n3. price (coin) - to show coin price\n4. exchange (currency val) (from which coin) (to which coin) to get current exchanges\nAvailable coins: usd, gst, sol\n"
}

func lazy() string {
	// show price
	sol := crypto.GetCoinData(crypto.SOL).MarketData.CurrentPrice["usd"]     // 1 sol = ? usd
	gst := crypto.GetCoinData(crypto.GST_SOL).MarketData.CurrentPrice["usd"] // 1 gst = ? usd

	ex := gst * 100 / sol // 1 gst = ? sol

	return fmt.Sprintf("1 sol = %.5f usd\n1 gst = %.5f usd\n100 gst = %.5f sol\n", sol, gst, ex)
}

func price(id string) string {
	// show price
	return fmt.Sprintf("1 %s = %.5f usd\n", id, crypto.GetCoinData(crypto.CoinNickname[id]).MarketData.CurrentPrice["usd"])
}

func exchange(count float64, from string, to string) string {
	f := crypto.GetCoinData(crypto.CoinNickname[from]).MarketData.CurrentPrice["usd"]
	t := crypto.GetCoinData(crypto.CoinNickname[to]).MarketData.CurrentPrice["usd"]

	return fmt.Sprintf("%.5f %s = %.5f %s\n", count, from, count*f/t, to)
}
