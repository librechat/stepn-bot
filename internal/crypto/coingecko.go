package crypto

import (
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

func Get() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	CG := coingecko.NewClient(httpClient)

	//CG.SimpleSinglePrice("", "")
	CG.CoinsList()
}
