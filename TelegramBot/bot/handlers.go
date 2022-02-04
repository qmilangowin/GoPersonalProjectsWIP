package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const coinGeckoAPI = "https://api.coingecko.com/api/v3/coins/"

func (app *BotApplication) home(w http.ResponseWriter, r *http.Request) {
	// url = "https://api.coingecko.com/api/v3/coins/quant-network"
	url := coinGeckoAPI + "quant-network"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var result TokenData
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}

	//TODO: REMOVE
	fmt.Println(result.ID)
	fmt.Println(result.MarketData.CurrentPrice.Usd)

	type Coin struct {
		Name float64 `json:"usd"`
	}

	var m Coin

	m.Name = result.MarketData.CurrentPrice.Usd

	if err := json.NewEncoder(w).Encode(&m); err != nil {
		log.Println(err)
		return
	}
}
