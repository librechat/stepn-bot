package handler

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/librechat/stepn-bot/internal/crypto"
)

var re = regexp.MustCompile(`([a-z0-9]*) ([a-z0-9]*) ([a-z0-9]*) ([a-z0-9]*)`)

func HandleCmd(line string) string {
	r := regex(line)
	switch r[1] {
	case "help":
		return help()
	case "price":
		return price()
	case "exchange":
		cnt, _ := strconv.ParseFloat(r[2], 64)
		return exchange(cnt, r[3], r[4])
	}
	return ""
}

func regex(line string) []string {
	return re.FindStringSubmatch(line)
}

func help() string {
	return "You may use below commands :-)\n1. help\n2. price - to show GST and SOL prices\n3. exchange (currency val) (from which coin) (to which coin) to get current exchanges\n"
}

func price() string {
	// show price
	sol := crypto.GetCoinData(crypto.SOL).MarketData.CurrentPrice["usd"]     // 1 sol = ? usd
	gst := crypto.GetCoinData(crypto.GST_SOL).MarketData.CurrentPrice["usd"] // 1 gst = ? usd

	ex := gst * 100 / sol // 1 gst = ? sol

	return fmt.Sprintf("1 SOL = %.5f USD\n1 GST = %.5f USD\n100 GST = %.5f SOL\n", sol, gst, ex)
}

func exchange(count float64, from string, to string) string {
	f := crypto.GetCoinData(from).MarketData.CurrentPrice["usd"]
	t := crypto.GetCoinData(to).MarketData.CurrentPrice["usd"]

	return fmt.Sprintf("%.5f %s = %.5f %s\n", count, from, count*f/t, to)
}
