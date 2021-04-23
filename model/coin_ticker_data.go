package model

import "time"

type CoinTickerData struct {
	Name    string    `json:"name"`
	Tickers []Tickers `json:"tickers"`
}
type Market struct {
	Name                string `json:"name"`
	Identifier          string `json:"identifier"`
	HasTradingIncentive bool   `json:"has_trading_incentive"`
}
type ConvertedLast struct {
	Btc float64 `json:"btc"`
	Eth float64 `json:"eth"`
	Usd float64 `json:"usd"`
}
type ConvertedVolume struct {
	Btc float64 `json:"btc"`
	Eth int     `json:"eth"`
	Usd int     `json:"usd"`
}
type Tickers struct {
	Base                   string          `json:"base"`
	Target                 string          `json:"target"`
	Market                 Market          `json:"market"`
	Last                   float64         `json:"last"`
	Volume                 float64         `json:"volume"`
	ConvertedLast          ConvertedLast   `json:"converted_last"`
	ConvertedVolume        ConvertedVolume `json:"converted_volume"`
	TrustScore             string          `json:"trust_score"`
	BidAskSpreadPercentage float64         `json:"bid_ask_spread_percentage"`
	Timestamp              time.Time       `json:"timestamp"`
	LastTradedAt           time.Time       `json:"last_traded_at"`
	LastFetchAt            time.Time       `json:"last_fetch_at"`
	IsAnomaly              bool            `json:"is_anomaly"`
	IsStale                bool            `json:"is_stale"`
	TradeURL               string          `json:"trade_url"`
	TokenInfoURL           interface{}     `json:"token_info_url"`
	CoinID                 string          `json:"coin_id"`
	TargetCoinID           string          `json:"target_coin_id,omitempty"`
}
