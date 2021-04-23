package api

import (
	"alert/crypto/model"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func MarketPriceByIds(ids string) model.MarketCurrencyData {
	url := "https://api.coingecko.com/api/v3/coins/markets"
	method := "GET"
	client := &http.Client {
		Timeout: time.Second * 15,
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	q := req.URL.Query()
	q.Add("vs_currency", "thb")
	if ids != "" {
		q.Add("ids", ids)
	}
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer res.Body.Close()

	var marketCurrencyData model.MarketCurrencyData
	_ = json.NewDecoder(res.Body).Decode(&marketCurrencyData)
	return marketCurrencyData
}
