package handler

import (
	"fmt"
	"strconv"

	"github.com/librechat/stepn-bot/internal/crypto"
)

func help(params ...string) string {
	return "Hi you may use below commands :-)\n1. help\n2. lazy - to show GST and SOL prices\n3. price (coin) - to show coin price\n4. exchange (currency val) (from which coin) (to which coin) to get current exchanges\nAvailable coins: usd, gst, sol\n"
}

func lazy(params ...string) string {
	// show price
	sol := crypto.GetCoinData(crypto.SOL).MarketData.CurrentPrice["usd"]     // 1 sol = ? usd
	gst := crypto.GetCoinData(crypto.GST_SOL).MarketData.CurrentPrice["usd"] // 1 gst = ? usd
	//gmt := crypto.GetCoinData(crypto.GMT).MarketData.CurrentPrice["usd"]
	//gstGmt := gst / gmt

	gstSol := gst * 100 / sol // 1 gst = ? sol

	return fmt.Sprintf("1 gst = %.5f usd\n1 sol = %.5f usd\n100 gst = %.5f sol\n", gst, sol, gstSol)
}

func richPrice(params ...string) string {
	if len(params) < 1 {
		return "Please input price (coin)"
	}

	target := params[0]
	exchanges, err := crypto.GetCoinRichExchange(target, crypto.RichExchangeSupported)
	if err != nil {
		return err.Error()
	}
	rate := int64(1)
	if len(params) > 1 {
		if rate, err = strconv.ParseInt(params[2], 10, 64); err != nil {
			return "Failed to parse exchange amount"
		}
	}

	msg := ""
	for i, coin := range crypto.RichExchangeSupported {
		if target != coin {
			msg += fmt.Sprintf("%d %s = %.5f %s\n", rate, target, float64(rate)*exchanges[i], coin)
		}
	}
	return msg
}

func price(params ...string) string {
	if len(params) < 1 {
		return "Please input price (coin)"
	}

	// show price
	p, err := crypto.GetCoinUSDPrice(params[0])
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("1 %s = %.5f usd\n", params[0], p)
}

func exchange(params ...string) string {
	if len(params) < 3 {
		return "Please input exchange (currency val) (from which coin) (to which coin)"
	}

	count, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return "Failed to read currency val"
	}

	ex, err := crypto.Exchange(count, params[1], params[2])
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%.5f %s = %.5f %s\n", count, params[1], ex, params[2])
}
