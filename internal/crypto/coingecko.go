package crypto

import (
	"log"
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
	coingecko "github.com/superoo7/go-gecko/v3"
	"github.com/superoo7/go-gecko/v3/types"
)

const (
	USDC    = "usd-coin"
	SOL     = "solana"
	GST_SOL = "green-satoshi-token"
)

var CoinNickname = map[string]string{
	"gst": GST_SOL,
	"sol": SOL,
	"usd": "usd",
}

// cache the price
var c = cache.New(5*time.Second, 10*time.Second)

func GetCoinData(id string) *types.CoinsID {
	if val, ok := c.Get(id); ok {
		return val.(*types.CoinsID)
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	coin, err := CG.CoinsID(id, false, true, true, true, true, false)
	if err != nil {
		log.Println(err.Error())
	}
	//fmt.Printf("%v, %v\n", coin.Name, coin.MarketData.CurrentPrice["usd"])
	c.Set(id, coin, cache.DefaultExpiration)
	return coin
}
