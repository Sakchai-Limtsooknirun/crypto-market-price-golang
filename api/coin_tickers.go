package api

import (
	"alert/crypto/model"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func CoinTickers() model.CoinTickerData {
	url := "https://api.coingecko.com/api/v3/coins/zipmex-token/tickers?exchange_ids=zipmex"
	method := "GET"
	client := &http.Client {
		Timeout: time.Second * 15,
	}
	req, err := http.NewRequest(method, url, nil)
	//url = strings.Replace(url, "{{id}}", "zipmex-token", 1)
	log.Println(url)
	if err != nil {
		log.Fatal(err)
		return model.CoinTickerData{}
	}
	log.Println(url+"2")

	//q := req.URL.Query()
	//q.Add("exchange_ids", "zipmex")
	//req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return model.CoinTickerData{}
	}
	log.Println(url+"3")

	defer res.Body.Close()

	var coinTickerData model.CoinTickerData
	_ = json.NewDecoder(res.Body).Decode(&coinTickerData)
	log.Println(url+"4")
	log.Println(coinTickerData.Tickers[0])
	return coinTickerData
}
